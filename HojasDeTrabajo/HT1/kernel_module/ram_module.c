#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/mm.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Xhunik Miguel");
MODULE_DESCRIPTION("Un módulo para obtener el uso de la memoria RAM");

static struct proc_dir_entry *archivo_proc;

// Función que se llama cuando se lee el archivo proc
static int mostrar_memoria(struct seq_file *m, void *v) {
    struct sysinfo i;
    si_meminfo(&i);
    seq_printf(m, "{\"Libre\": %lu,", i.freeram * 4);
    seq_printf(m, "\"Total\": %lu", i.totalram * 4);
    seq_printf(m, "}");
    return 0;
}

static int abrir_archivo_proc(struct inode *inode, struct file *file) {
    return single_open(file, mostrar_memoria, NULL);
}

static const struct proc_ops ops_archivo_proc = {
    .proc_open = abrir_archivo_proc,
    .proc_read = seq_read,

};

static int __init iniciar_modulo(void) {
    archivo_proc = proc_create("info_memoria", 0, NULL, &ops_archivo_proc);
    if (!archivo_proc) {
        return -ENOMEM;
    }
    return 0;
}

static void __exit limpiar_modulo(void) {
    proc_remove(archivo_proc);
}

module_init(iniciar_modulo);
module_exit(limpiar_modulo);
