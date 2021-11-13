const Web3 = require("web3");
let eth = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:8545/")).eth;
const start = async function () {
    do {
        let account = eth.accounts.create()
        await eth.getBalance(account.address).then(
            (data) => {
                if (data != "0") {
                    console.log("%s: %s", account, data);
                }
            }
        );
    } while (true);
}
start();