use sysinfo::System;
use colored::*;
use std::collections::HashSet;

pub struct ProcessDetails {
    pub pid: u32,
    pub name: String,
}

pub fn get_process_details(pids: &[u32]) -> Vec<ProcessDetails> {
    let mut sys = System::new_all();
    sys.refresh_all();

    let mut details = Vec::new();
    let pid_set: HashSet<u32> = pids.iter().cloned().collect();

    for (pid, process) in sys.processes() {
        let pid_u32 = pid.as_u32();
        if pid_set.contains(&pid_u32) {
            details.push(ProcessDetails {
                pid: pid_u32,
                name: process.name().to_string_lossy().to_string(),
            });
        }
    }
    details
}

pub fn print_process_table(details: &[ProcessDetails]) {
    if details.is_empty() {
        println!("{}", "No locking processes found.".green());
        return;
    }

    println!("\n{}", "Locking Processes Found:".yellow().bold());
    println!("{:<8} {:<25}", "PID".bold(), "Process Name".bold());
    println!("{}", "-".repeat(35));

    for detail in details {
        println!("{:<8} {:<25}", detail.pid, detail.name.cyan());
    }
    println!();
}

pub fn select_processes_to_kill(details: &[ProcessDetails]) -> Vec<u32> {
    if details.is_empty() {
        return Vec::new();
    }

    let options: Vec<String> = details.iter()
        .map(|d| format!("{} ({})", d.name, d.pid))
        .collect();

    let ans = inquire::MultiSelect::new("Select processes to terminate:", options).prompt();

    match ans {
        Ok(selected) => {
            details.iter()
                .filter(|d| {
                    let formatted = format!("{} ({})", d.name, d.pid);
                    selected.contains(&formatted)
                })
                .map(|d| d.pid)
                .collect()
        }
        Err(_) => Vec::new(),
    }
}
