
const express = require("express");
const bodyParser = require("body-parser");

const app = express();
app.use(bodyParser.json());

app.post("/createAsset", async (req, res) => {
    const { id, dealerId, msisdn, mpin, balance, status, transAmount, transType, remarks } = req.body;
    // simulate Fabric call
    return res.json({ success: true, message: "Asset created (simulated)." });
});

app.get("/readAsset/:id", async (req, res) => {
    const { id } = req.params;
    // simulate Fabric call
    return res.json({ id, dealerId: "D001", msisdn: "1234567890" });
});

app.listen(3000, () => {
    console.log("API server running on port 3000");
});
