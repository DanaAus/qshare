mod locker;
mod cli;
mod ui;

use clap::Parser;
use cli::Cli;
use std::process;
use colored::*;

fn main() {
    let cli = Cli::parse();

    if !cli.path.exists() {
        eprintln!("{} Path does not exist: {:?}", "Error:".red().bold(), cli.path);
        process::exit(1);
    }

    println!("{} Scanning for locks on {:?}...", "Info:".blue().bold(), cli.path);

    let pids = locker::get_locking_processes(&cli.path);
    
    if pids.is_empty() {
        println!("{}", "No locking processes found.".green().bold());
        return;
    }

    let details = ui::get_process_details(&pids);
    ui::print_process_table(&details);

    if cli.force {
        println!("{}", "Force kill enabled. Terminating all locking processes...".red().bold());
        // TODO: Implement kill logic in Phase 4
    } else {
        let _to_kill = ui::select_processes_to_kill(&details);
        // TODO: Implement kill logic in Phase 4
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
