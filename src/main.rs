#![feature(proc_macro_hygiene, decl_macro)]


#[macro_use] extern crate rocket;

// use serde::Deserialize;
// use rocket_contrib::json::Json;

use rocket::State;
use rocket::Data;

use std::sync::Mutex;

// extern crate serde_json;

// use serde_json::{Map, Value};

use std::collections::HashMap;


/*
#[derive(Deserialize)]
struct Folder {
    name: String,
    owner: String,
    // access_rights: AccessRights,
    allowed: Vec<String>,
    files: Vec<File>,
    folders: Vec<Folder>
}

impl Folder {
    fn new(name: &str, owner: &str, )
}

#[derive(Deserialize)]
struct File {
    name: String,
    owner: String,
    // access_rights: AccessRights,
    allowed: Vec<String>,
    data: Vec<u8>
}

struct Container {
    root: Folder
}

impl Container {
    fn new() -> Self {
        Self {root: Folder::new("$", )}
    }

    fn root() -> Self {
        Self {name: "/".to_string(), allowed: vec![], data: None, children: None}
    }

    push_child(contianer: Container) {
        if let Some(mut children) = self.children {
            children.push(contianer)
        } else {
            self.children = Some(vec![container])
        }
    }

    // remove_child(name: &str)

    replace_child(&mut self, contianer: Contianer){
        // if exists remove and push
        // else push
    }

    get_child(&self, name: &str) -> Container {
        if let Some(children) = self.children {
            children.iter().find(|x| x.name == name)
        }
        
    }

    // get_childs_value(name: &str) -> Container {

    // }

}

*/

struct SecretStorage {
    storage: Mutex<HashMap<String, Vec<u8>>>
}

impl SecretStorage {
    fn new() -> Self {
        Self{storage: Mutex::new(HashMap::new())}
    }

    // fn add_s(&mut self, key: &str, value: &str) -> {

    // }

    fn add(&self, key: &str, value: Vec<u8>) { // &[u8]
        let mut lock = self.storage.lock().unwrap();
        lock.insert(key.to_string(), value.to_vec());
    }

    fn get(&self, key: &str) -> Option<Vec<u8>> {
        let lock = self.storage.lock().unwrap();
        let value = lock.get(key);
        value.map(|x|x.to_vec())
    }
}

#[post("/<path>", format = "plain", data = "<data>")]
fn store(path: String, ss: State<SecretStorage>, data: Data) -> String {
    use std::io::{Cursor, Read, Seek, SeekFrom, Write};

    // Create fake "file"
    let mut c = Cursor::new(Vec::new());

    // Write into the "file" and seek to the beginning
    // c.write_all(&[1, 2, 3, 4, 5]).unwrap();
    data.stream_to(&mut c).map(|n| format!("Wrote {} bytes.", n));
    c.seek(SeekFrom::Start(0)).unwrap();

    // Read the "file's" contents into a vector
    let mut out = Vec::new();
    c.read_to_end(&mut out).unwrap();
    ss.add(&path, out);
    format!("Post {}", path)
}

#[get("/<path>")]
fn fetch(path: String, ss: State<SecretStorage>) -> Option<String> {
    if let Some(value) = ss.get(&path){
        Some(format!("Get {}", std::str::from_utf8(&value).unwrap()))
    } else {
        None
    }
    
}

fn main() {
    rocket::ignite()
    .manage(SecretStorage::new())
    .mount("/storage", routes![store, fetch]).launch();
}

// it receives a server public user specific encrypted date specific and request

// Server identifies the user using date, decodes request and send request with 