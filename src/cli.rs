use clap::Parser;
use std::path::PathBuf;

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
pub struct Cli {
    /// The file or directory to unlock
    pub path: PathBuf,

    /// Force kill the processes
    #[arg(short, long)]
    pub force: bool,

    /// Recursive scan
    #[arg(short, long, default_value_t = true)]
    pub recursive: bool,
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_cli_parsing_basic() {
        let args = vec!["fileinbreach", "test.txt"];
        let cli = Cli::try_parse_from(args).unwrap();
        assert_eq!(cli.path, PathBuf::from("test.txt"));
        assert!(!cli.force);
        assert!(cli.recursive);
    }

    #[test]
    fn test_cli_parsing_force() {
        let args = vec!["fileinbreach", "test.txt", "--force"];
        let cli = Cli::try_parse_from(args).unwrap();
        assert!(cli.force);
    }
}
