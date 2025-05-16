use std::io::Read;
use blake3;
const CHUNK_SIZE: usize = 1024 * 1024; // 1MB

fn chunk_file(path: &str) -> std::io::Result<Vec<Vec<u8>>> {
  let mut file = std::fs::File::open(path)?;
  let mut chunks = Vec::new();

  loop {
    let mut buffer = vec![0; CHUNK_SIZE];
    let n = file.read(&mut buffer)?;
    if n == 0 { break; }
    buffer.truncate(n);
    chunks.push(buffer);
  }

  Ok(chunks)
}

fn calc_hash(data: &[u8]) -> String {
  blake3::hash(data).to_hex().to_string()
}

fn store_metadata(db: &sled::Db, file_id: &str, chunk_hashes: &[String]) -> sled::Result<()> {
  let data = serde_json::to_vec(chunk_hashes).unwrap();
  db.insert(file_id, data)?;
  Ok(())
}