#![feature(proc_macro_hygiene, decl_macro)]


#[macro_use] extern crate rocket;

use rocket::State;

extern crate serde_json;

// use serde_json::{Map, Value};

use std::collections::HashMap;

struct SecretStorage {
    storage: HashMap<String, Vec<u8>>
}

impl SecretStorage {
    fn new() -> Self {
        Self{storage: HashMap::new()}
    }

    // fn add_s(&mut self, key: &str, value: &str) -> {

    // }

    fn add(&mut self, key: &str, value: &[u8]) {
        self.storage.insert(key.to_string(), value.to_vec());
    }

    fn get(&mut self, key: &str) -> Option<&Vec<u8>> {
        self.storage.get(key)
    }
}

#[post("/<path>")]
fn store(path: String, ss: State<SecretStorage>) -> String {
    ss.add(&path, &vec![240, 159, 146, 150]);
    format!("Post {}", path)
}

#[get("/<path>")]
fn fetch(path: String, ss: State<SecretStorage>) -> String {
    let value = ss.get(&path).unwrap();
    format!("Get {}", std::str::from_utf8(value).unwrap())
}

fn main() {
    rocket::ignite()
    .manage(SecretStorage::new())
    .mount("/api", routes![store, fetch]).launch();
}