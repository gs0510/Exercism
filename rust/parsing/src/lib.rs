use std::fs::File;
use std::io::{BufRead, BufReader, Result};

struct Alias {
    machine_name: String,
    host_name: String,
    machine_number: String,
}

struct MachineRoutingInfo {
    machine_name: String,
    machine_number: String,
    machine_type: String,
    machine_level: String,
    machine_location: String,
    machine_pods: ,
    machine_owner: String,
}

fn tokenizer() -> {


}

fn readAlias() -> std::io::Result<()>{
    let mut aliases = HashSet::new();
    let mut f = File::open("bbcpu.alias")?;
    for line in BufReader::new(file).lines() {
        tokenized_string = tokenizer(line)
        aliases.insert(Alias{machine_name: tokenized_string[0], host_name: tokenized_string[1], machine_number: tokenized_string[2]})
    }
    Ok(())
}

fn readList() -> {

}

fn main() -> std::io::Result<()> {
    let mut f = File::open("foo.txt")?;
    Ok(())
}

