"use strict";

const { Gateway, Wallets } = require("fabric-network");
const yaml = require("js-yaml");
const fs = require("fs");
const path = require("path");
const mspId = "Org2MSP";
const CC_NAME = "monitoring";
const CHANNEL = "mychannel";
let ccp = null;


/* Contoh Payload createAloptama
{
    channelName:"mychannel",
    userId:"user",
    data:{
        kodealat:"03072200",
        namaalat:"nodeJS",
        merekalat:"Intel Xeon",
        jumlahalat:1,
        tahunpengadaan:2022,
        kondisi:"Baru",
        keterangan:"OK"
    }
}
*/

exports.createAloptama = async (request) => {
        // load the network configuration
        const ccpPath = path.resolve(
            __dirname,
            "connection-org.yaml"
        );
        if (ccpPath.includes(".yaml")) {
            ccp = yaml.load(fs.readFileSync(ccpPath, "utf8"));
        } else {
            ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));
        }

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), "wallet", mspId);
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(request.userId);
        if (!identity) {
            console.log(
                'An identity for the user "${userId}" does not exist in the wallet'
            );
            console.log("Run the registerUser.js application before retrying");
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, {
            wallet,
            identity: request.userId,
            discovery: { enabled: true, asLocalhost: false },
        });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(request.channelName);

        // Get the contract from the network.
        const contract = network.getContract(CC_NAME);

        let data = request.data;
        let result = await contract.submitTransaction('CreateAloptama', data.kodealat, data.namaalat, data.merekalat, data.jumlahalat, data.tahunpengadaan, data.kondisi, data.keterangan);
        console.log("Transaction has been submitted");

        return (result);
}

/* Contoh Payload createAlatoto
{
    channelName:"mychannel",
    userId:"user",
    data:{
        kodesite:"03072200",
        namasite:"ARG Test 123",
        jenisalat:"ARG",
        lokasisite:"STMKG",
        mereksensor:"Vaisala",
        mereklogger:"Vaisala",
        ressensor:0.2,
        kapbaterai:40,
        kapsolar:50,
        corrmt:"",
        prevmt:""
    }
}
*/

exports.createAlatoto = async (request) => {
    // load the network configuration
    const ccpPath = path.resolve(
        __dirname,
        "connection-org.yaml"
    );
    if (ccpPath.includes(".yaml")) {
        ccp = yaml.load(fs.readFileSync(ccpPath, "utf8"));
    } else {
        ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));
    }

    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), "wallet", mspId);
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    // Check to see if we've already enrolled the user.
    const identity = await wallet.get(request.userId);
    if (!identity) {
        console.log(
            'An identity for the user "${userId}" does not exist in the wallet'
        );
        console.log("Run the registerUser.js application before retrying");
        return;
    }

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, {
        wallet,
        identity: request.userId,
        discovery: { enabled: true, asLocalhost: false },
    });

    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork(request.channelName);

    // Get the contract from the network.
    const contract = network.getContract(CC_NAME);

    let data = request.data;
    let result = await contract.submitTransaction('CreateAlatOto', data.kodesite, data.namasite, data.jenisalat, data.lokasisite, data.mereksensor, data.mereklogger, data.ressensor, data.kapbaterai, data.kapsolar, data.corrmt, data.prevmt);
    console.log("Transaction has been submitted");

    return (result);
}

/* Contoh Payload updateKondisiAloptama
{
    channelName:"mychannel",
    userId:"user",
    data:{
        kodealat:"03072200",
        newKondisi:"Update API",
        newKeterangan:"Update Berhasil"
    }
}
*/

exports.updateKondisiAloptama = async (request) => {
    // load the network configuration
    const ccpPath = path.resolve(
        __dirname,
        "connection-org.yaml"
    );
    if (ccpPath.includes(".yaml")) {
        ccp = yaml.load(fs.readFileSync(ccpPath, "utf8"));
    } else {
        ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));
    }

    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), "wallet", mspId);
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    // Check to see if we've already enrolled the user.
    const identity = await wallet.get(request.userId);
    if (!identity) {
        console.log(
            'An identity for the user "${userId}" does not exist in the wallet'
        );
        console.log("Run the registerUser.js application before retrying");
        return;
    }

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, {
        wallet,
        identity: request.userId,
        discovery: { enabled: true, asLocalhost: false },
    });

    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork(request.channelName);

    // Get the contract from the network.
    const contract = network.getContract(CC_NAME);

    let data = request.data;
    let result = await contract.submitTransaction('UpdateKondisiAloptama', data.kodealat, data.newKondisi, data.newKeterangan);
    console.log("Transaction has been submitted");

    return (result);
}

/* Contoh Payload updatePMCM
{
    channelName:"mychannel",
    userId:"user",
    data:{
        kodealat:"03072200",
        newPrevMT:"Ganti AKi",
        newCorrMT:""
    }
}
*/

exports.updatePMCM = async (request) => {
    // load the network configuration
    const ccpPath = path.resolve(
        __dirname,
        "connection-org.yaml"
    );
    if (ccpPath.includes(".yaml")) {
        ccp = yaml.load(fs.readFileSync(ccpPath, "utf8"));
    } else {
        ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));
    }

    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), "wallet", mspId);
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    // Check to see if we've already enrolled the user.
    const identity = await wallet.get(request.userId);
    if (!identity) {
        console.log(
            'An identity for the user "${userId}" does not exist in the wallet'
        );
        console.log("Run the registerUser.js application before retrying");
        return;
    }

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, {
        wallet,
        identity: request.userId,
        discovery: { enabled: true, asLocalhost: false },
    });

    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork(request.channelName);

    // Get the contract from the network.
    const contract = network.getContract(CC_NAME);

    let data = request.data;
    let result = await contract.submitTransaction('UpdatePMCM', data.kodesite, data.newPrevMT, data.newCorrMT);
    console.log("Transaction has been submitted");

    return (result);
}


/* async function invoke(user) {
    try {
        console.log("Invoking chaincode using : ", user);
        // load the network configuration
        const ccpPath = path.resolve(
            __dirname,
            "connection-org.yaml"
        );
        if (ccpPath.includes(".yaml")) {
            ccp = yaml.load(fs.readFileSync(ccpPath, "utf8"));
        } else {
            ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));
        }
        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), "wallet", mspId);
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(user);
        if (!identity) {
            console.log(
                'An identity for the user "${user}" does not exist in the wallet'
            );
            console.log("Run the registerUser.js application before retrying");
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, {
            wallet,
            identity: user,
            discovery: { enabled: true, asLocalhost: false },
        });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(CHANNEL);

        // Get the contract from the network.
        const contract = network.getContract(CC_NAME);

        const result = await contract.submitTransaction(
            "CreateAloptama",
            "010030062022", "Node Server", "Dewacloud", '1', '2022', "Baik", "OK"
        );
        console.log("Transaction has been submitted");

        // Disconnect from the gateway.
        gateway.disconnect();
        return result;
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}
invoke("appUser") */