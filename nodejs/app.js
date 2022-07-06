"use strict"

const express = require("express");
const app = express();
var morgan = require('morgan')
app.use(morgan('combined'))
const bodyparser = require("body-parser");

var cors = require('cors');
const {registerUser, userExist} = require("./app/registerUser");
const {createAloptama, createAlatoto, updateKondisiAloptama, updatePMCM} = require("./app/invoke");
app.use(cors())
app.use(bodyparser.json());

app.listen(4000, () => {
    console.log("Server running on port 4000");
})

app.post("/register", async (req,res) => {
    try {
        let org = req.body.org;
        let userId = req.body.userId;
        let result = await registerUser({ OrgMSP: org, userId: userId });
        res.send(result);
    } catch (error) {
        res.status(500).send(error)
    }

});

app.post("/createAloptama", async (req,res) => {
    try {
        let payload = {
            "channelName": req.body.channelName,
            "userId": req.body.userId,
            "data": req.body.data
        }

        let result = await createAloptama(payload);
        res.send(result)
    } catch(error) {
        res.status(500).send(error)
    }   
});

app.post("/createAlatoto", async (req,res) => {
    try {
        let payload = {
            "channelName": req.body.channelName,
            "userId": req.body.userId,
            "data": req.body.data
        }

        let result = await createAlatoto(payload);
        res.send(result)
    } catch(error) {
        res.status(500).send(error)
    }
});

app.post("/updateKondisiAloptama", async (req,res) => {
    try {
        let payload = {
            "channelName": req.body.channelName,
            "userId": req.body.userId,
            "data": req.body.data
        }

        let result = await updateKondisiAloptama(payload);
        res.send(result)
    } catch(error) {
        res.status(500).send(error)
    }
});

app.post("/updatePMCM", async (req,res) => {
    try {
        let payload = {
            "channelName": req.body.channelName,
            "userId": req.body.userId,
            "data": req.body.data
        }

        let result = await updatePMCM(payload);
        res.send(result)
    } catch(error) {
        res.status(500).send(error)
    }
});