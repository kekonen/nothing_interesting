const { MongoClient } = require("mongodb");
// Connection URI
const uri = "mongodb://localhost:27017";
// Create a new MongoClient


(async () => {
    const client = new MongoClient(uri);
    await client.connect();
    // Establish and verify connection
    const db = client.db("admin");
    await db.collection('keks').insertMany([
        {lol:1, kek:2},
        {lol:3, kek:4},
    ])

    const result = await db.collection('keks').find({kek:2});
    result.forEach(x=>{console.log(`xxx`,x)})
    db.collection('keks').deleteMany({kek:{
        $or: [2,4]
    }})

    console.log("Connected successfully to server");
    await client.close()
})()