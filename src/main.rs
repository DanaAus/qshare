mod locker;
mod cli;
mod ui;
mod killer;
mod logger;

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

    let start_time = std::time::Instant::now();
    let pids = locker::get_locking_processes(&cli.path);
    let duration = start_time.elapsed();
    
    if pids.is_empty() {
        println!("{} (took {:?})", "No locking processes found.".green().bold(), duration);
        return;
    }

    println!("{} Found {} locking processes in {:?}.", "Info:".blue().bold(), pids.len(), duration);
    let details = ui::get_process_details(&pids);
    ui::print_process_table(&details);

    let to_kill = if cli.force {
        println!("{}", "Force kill enabled. Targeting all locking processes...".red().bold());
        pids
    } else {
        ui::select_processes_to_kill(&details)
    };

    if to_kill.is_empty() {
        println!("{}", "No processes selected for termination.".yellow());
        return;
    }

    logger::log_action(&format!("Scan complete for {:?}. Found {} processes.", cli.path, details.len()));

    for pid in to_kill {
        let detail = details.iter().find(|d| d.pid == pid);
        let name = detail.map(|d| d.name.as_str()).unwrap_or("Unknown");

        if killer::is_critical_process(name) {
            println!("{} Process {} ({}) is a critical system process!", "Warning:".yellow().bold(), name, pid);
            let confirm = inquire::Confirm::new("Are you absolutely sure you want to terminate it?")
                .with_default(false)
                .prompt();
            
            if let Ok(false) | Err(_) = confirm {
                println!("Skipping critical process {} ({}).", name, pid);
                logger::log_action(&format!("Skipped critical process {} ({})", name, pid));
                continue;
            }
        }

        println!("{} Terminating process {} ({})...", "Action:".magenta().bold(), name, pid);
        match killer::kill_process(pid, cli.force) {
            Ok(_) => {
                println!("{} Successfully terminated {} ({}).", "Success:".green(), name, pid);
                logger::log_action(&format!("Terminated process {} ({})", name, pid));
            }
            Err(e) => {
                eprintln!("{} Failed to terminate {} ({}): {}", "Error:".red(), name, pid, e);
                logger::log_action(&format!("Failed to terminate process {} ({}): {}", name, pid, e));
            }
        }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
