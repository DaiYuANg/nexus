use clap::Parser;

/// Simple program to greet a person
#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
pub(crate) struct Args {
  #[arg(short, long, default_value = "config.toml")]
  config: String,

  #[arg(short, long)]
  debug: bool,
}
