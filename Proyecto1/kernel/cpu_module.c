#include <linux/module.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/sched/signal.h>
#include <linux/seq_file.h>
#include <linux/fs.h>
#include <linux/sched.h>
#include <linux/mm.h>
#include <linux/delay.h>
#include <linux/sched/loadavg.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Xhunik Miguel");
MODULE_DESCRIPTION("Descriptor de CPU");
MODULE_VERSION("1.0");

static int write_to_proc(struct seq_file *m, void *v);
static int open_proc(struct inode *inode, struct file *file);

static struct proc_ops file_ops = {
    .proc_open = open_proc,
    .proc_read = seq_read
};

static int __init module_start(void)
{
    proc_create("cpu_so1_1s2024", 0, NULL, &file_ops);
    printk(KERN_INFO "CPU Info Module Loaded\n");
    return 0;
}

static void __exit module_end(void)
{
    remove_proc_entry("cpu_so1_1s2024", NULL);
    printk(KERN_INFO "CPU Info Module Removed\n");
}

module_init(module_start);
module_exit(module_end);

static int write_to_proc(struct seq_file *m, void *v) {
    struct task_struct *task, *child_task;
    struct list_head *list;
    unsigned long rss = 0;
    int running = 0, sleeping = 0, zombie = 0, stopped = 0;

    unsigned long total_cpu_usage = 0;
    total_cpu_usage = (avenrun[0] << 16) + (avenrun[1] >> 16);

    
    seq_printf(m, "{\n\"Total_CPU_Time\":%lu,\n", total_cpu_usage);

    seq_printf(m, "\"Processes\":[\n");

    int first_entry = 1;
    for_each_process(task) {
        rss = task->mm ? get_mm_rss(task->mm) << PAGE_SHIFT : 0;

        seq_printf(m, "%s\n{\"PID\":%d,\n", first_entry ? "" : ",", task->pid);
        first_entry = 0;
        seq_printf(m, "\"Name\":\"%s\",\n", task->comm);
        seq_printf(m, "\"User\": %u,\n", from_kuid(&init_user_ns, task->cred->uid));
        seq_printf(m, "\"State\":%u,\n", task->__state);
        seq_printf(m, "\"RAM_Percentage\":%lu,\n", (rss * 100) / totalram_pages());
        seq_printf(m, "\"RSS\":%lu,\n", rss);

        seq_printf(m, "\"Children\":[\n");
        int first_child = 1;
        list_for_each(list, &(task->children)) {
            child_task = list_entry(list, struct task_struct, sibling);

            rss = child_task->mm ? get_mm_rss(child_task->mm) << PAGE_SHIFT : 0;
            seq_printf(m, "%s{\"PID\":%d,\n", first_child ? "" : ",", child_task->pid);
            first_child = 0;
            seq_printf(m, "\"Name\":\"%s\",\n", child_task->comm);
            seq_printf(m, "\"State\":%u,\n", child_task->__state);
            seq_printf(m, "\"Parent_PID\":%d,\n", task->pid);
            seq_printf(m, "\"RSS\":%lu}\n", rss);
        }
        seq_printf(m, "]}\n");

        switch (task->__state) {
            case 0: running++; break;
            case 1: sleeping++; break;
            case 4: zombie++; break;
            default: stopped++;
        }
    }

    seq_printf(m, "],\n");
    seq_printf(m, "\"Running\":%d,\n\"Sleeping\":%d,\n\"Zombie\":%d,\n\"Stopped\":%d,\n\"Total\":%d\n}", 
               running, sleeping, zombie, stopped, running + sleeping + zombie + stopped);

    return 0;
}

static int open_proc(struct inode *inode, struct file *file) {
    return single_open(file, write_to_proc, NULL);
}
