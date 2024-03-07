#include <linux/module.h>
#include <linux/proc_fs.h>
#include <linux/sysinfo.h>
#include <linux/seq_file.h>
#include <linux/mm.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Xhunik Miguel");
MODULE_DESCRIPTION("Descriptor de memoria RAM");

struct sysinfo info_mem;

static int show_ram_info(struct seq_file *m, void *v)
{
    unsigned long total_ram, used_ram, free_ram;
    unsigned long usage_percentage;
    si_meminfo(&info_mem);

    total_ram = info_mem.totalram * info_mem.mem_unit >> 20;
    free_ram = (info_mem.freeram * info_mem.mem_unit + info_mem.bufferram * info_mem.mem_unit + info_mem.sharedram * info_mem.mem_unit) >> 20;
    used_ram = total_ram - free_ram;
    usage_percentage = (used_ram * 100) / total_ram;
    
    seq_printf(m, "{\"TotalRAM\":%lu, \"UsedMemory\":%lu, \"UsagePercent\":%lu, \"FreeMemory\":%lu}", total_ram, used_ram, usage_percentage, free_ram);
    return 0;
}

static int open_ram_proc(struct inode *inode, struct file *file)
{
    return single_open(file, show_ram_info, NULL);
}

static struct proc_ops ram_proc_fops = {
    .proc_open = open_ram_proc,
    .proc_read = seq_read};

static int __init ram_module_init(void)
{
    proc_create("ram_so1_1s2024", 0, NULL, &ram_proc_fops);
    printk(KERN_INFO "RAM module loaded\n");
    return 0;
}

static void __exit ram_module_exit(void)
{
    remove_proc_entry("ram_so1_1s2024", NULL);
}

module_init(ram_module_init);
module_exit(ram_module_exit);
