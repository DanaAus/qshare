use std::path::{Path, PathBuf};
use std::os::windows::ffi::OsStrExt;
use windows::Win32::System::RestartManager::{
    RmStartSession, RmRegisterResources, RmGetList, RmEndSession,
    RM_PROCESS_INFO,
};
use windows::Win32::Foundation::{ERROR_SUCCESS, ERROR_MORE_DATA};
use windows::core::{PWSTR, PCWSTR};
use walkdir::WalkDir;
use std::collections::HashSet;

/// Identifies processes that are currently locking the given file or directory.
/// If a directory is provided, it recursively finds locks for all files within it.
pub fn get_locking_processes(path: &Path) -> Vec<u32> {
    let mut files = Vec::new();
    if path.is_file() {
        files.push(path.to_path_buf());
    } else if path.is_dir() {
        for entry in WalkDir::new(path).into_iter().filter_map(|e| e.ok()) {
            if entry.file_type().is_file() {
                files.push(entry.path().to_path_buf());
            }
        }
    } else {
        // Path doesn't exist or is something else
        return Vec::new();
    }

    if files.is_empty() {
        return Vec::new();
    }

    get_locking_processes_for_files(&files)
}

/// Identifies processes that are currently locking the given list of files.
pub fn get_locking_processes_for_files(files: &[PathBuf]) -> Vec<u32> {
    let mut session_handle = 0u32;
    let mut session_key = [0u16; 33];
    
    unsafe {
        let res = RmStartSession(&mut session_handle, Some(0), PWSTR(session_key.as_mut_ptr()));
        if res != ERROR_SUCCESS {
            return Vec::new();
        }

        // Collect all wide paths
        let paths_wide: Vec<Vec<u16>> = files.iter()
            .map(|p| p.as_os_str().encode_wide().chain(std::iter::once(0)).collect())
            .collect();
        
        let paths: Vec<PCWSTR> = paths_wide.iter().map(|pw| PCWSTR(pw.as_ptr())).collect();
        
        let res = RmRegisterResources(session_handle, Some(&paths), None, None);
        if res != ERROR_SUCCESS {
            let _ = RmEndSession(session_handle);
            return Vec::new();
        }

        let mut n_proc_info_needed = 0u32;
        let mut n_proc_info = 0u32;
        let mut reboot_reasons = 0u32;
        
        let res = RmGetList(
            session_handle,
            &mut n_proc_info_needed,
            &mut n_proc_info,
            None,
            &mut reboot_reasons,
        );

        if res != ERROR_SUCCESS && res != ERROR_MORE_DATA {
            let _ = RmEndSession(session_handle);
            return Vec::new();
        }

        if n_proc_info_needed == 0 {
            let _ = RmEndSession(session_handle);
            return Vec::new();
        }

        let mut process_info = vec![RM_PROCESS_INFO::default(); n_proc_info_needed as usize];
        n_proc_info = n_proc_info_needed;

        let res = RmGetList(
            session_handle,
            &mut n_proc_info_needed,
            &mut n_proc_info,
            Some(process_info.as_mut_ptr()),
            &mut reboot_reasons,
        );

        let mut pids = HashSet::new();
        if res == ERROR_SUCCESS {
            for i in 0..n_proc_info as usize {
                pids.insert(process_info[i].Process.dwProcessId);
            }
        }

        let _ = RmEndSession(session_handle);
        pids.into_iter().collect()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::fs::File;
    use tempfile::tempdir;

    #[test]
    fn test_get_locking_processes_none() {
        let dir = tempdir().unwrap();
        let file_path = dir.path().join("test.txt");
        File::create(&file_path).unwrap();

        let processes = get_locking_processes(&file_path);
        assert!(processes.is_empty());
    }

    #[test]
    fn test_get_locking_processes_exists() {
        let dir = tempdir().unwrap();
        let file_path = dir.path().join("locked.txt");
        File::create(&file_path).unwrap();
        
        let _file_handle = File::options().read(true).write(true).open(&file_path).unwrap();
        
        let processes = get_locking_processes(&file_path);
        assert!(!processes.is_empty());
        
        let current_pid = std::process::id();
        assert!(processes.contains(&current_pid));
    }

    #[test]
    fn test_get_locking_processes_recursive() {
        let dir = tempdir().unwrap();
        let sub_dir = dir.path().join("subdir");
        std::fs::create_dir(&sub_dir).unwrap();
        
        let file1_path = dir.path().join("file1.txt");
        let file2_path = sub_dir.join("file2.txt");
        
        File::create(&file1_path).unwrap();
        File::create(&file2_path).unwrap();
        
        let _handle1 = File::options().read(true).write(true).open(&file1_path).unwrap();
        let _handle2 = File::options().read(true).write(true).open(&file2_path).unwrap();
        
        let processes = get_locking_processes(dir.path());
        assert!(!processes.is_empty());
        
        let current_pid = std::process::id();
        assert!(processes.contains(&current_pid));
    }
}
