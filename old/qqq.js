/**
 * Created by zim on 19-Jan-18.
 */
Skip to content
This repository
Search
Pull requests
Issues
Marketplace
Explore
@yohoho-ua
Sign out
Unwatch 2
Star 0  Fork 1 yohoho-ua/Hippodrome-Dapp
Code  Issues 0  Pull requests 1  Projects 0  Wiki  Insights  Settings
Branch: master Find file Copy pathHippodrome-Dapp/app/javascripts/app.js
f0fa168  12 days ago
@yohoho-ua yohoho-ua some code cleaning
2 contributors @yohoho-ua @arhanechka
RawBlameHistory
151 lines (127 sloc)  5.3 KB
// Import the page's CSS. Webpack will know what to do with it.
import '../stylesheets/app.css'

// Import libraries we need.
import { default as Web3 } from 'web3'
import { default as contract } from 'truffle-contract'

// Import our contract artifacts and turn them into usable abstractions.
import hippodrome_artifacts from '../../build/contracts/Hippodrome.json'

// MetaCoin is our usable abstraction, which we'll use through the code below.
var Hippodrome = contract(hippodrome_artifacts)

// The following code is simple to show off interacting with your contracts.
// As your needs grow you will likely need to change its form and structure.
// For application bootstrapping, check out window.addEventListener below.
var accounts
var account
var hippo

window.App = {
    start: function () {
        var self = this

        // Bootstrap the Hippodrome abstraction for Use.
        Hippodrome.setProvider(web3.currentProvider)

        Hippodrome.deployed().then(function (instance) {
            hippo = instance
        })

        // Get the initial account so it can be displayed.
        web3.eth.getAccounts(function (err, accs) {
            if (err != null) {
                alert('There was an error fetching your accounts.')
                return
            }

            if (accs.length == 0) {
                alert("Couldn't get any accounts! Make sure your Ethereum client is configured correctly.")
                return
            }

            accounts = accs
            account = accounts[0]

            var accountInterval = setInterval(function() {
                if (web3.eth.defaultAccount!== account) {
                    account = web3.eth.accounts[0];
                    window.App.updateCurrentAcc();
                }
            }, 1000);
        })

        Hippodrome.deployed().then(function(deployed) {
            return deployed.totalBet();
        }).then(function (result) {
            console.log(result.c[0]);
        })

        Hippodrome.deployed().then(function(deployed) {
            return deployed.info();
        }).then(function (result) {
            console.log(result);
        })
    },

    updateStatus: function (message) {
        var status = document.getElementById('status')
        status.innerHTML = message
    },

    updateCurrentAcc: function () {
        var curr_acc_element = document.getElementById('acc')
        curr_acc_element.innerHTML = account.valueOf()
    },

    setMaxPlayers: function () {
        var self = this

        var amount = parseInt(document.getElementById('maxPlayersInput').value)
        this.setStatus('Initiating transaction... (please wait)')

        hippo.setMaxPlayers(amount, { from: account })
            .then(function (result) {
                self.updateStatus('Transaction complete!')
                self.updateAll(result)
            }).catch(function (e) {
            console.log(e)
            self.updateStatus('Error sending transaction; see log.')
        })
    },

    bet: function () {
        var self = this
        var horseNumber = parseInt(document.getElementById('horseNumber').value)
        this.setStatus('Initiating transaction... (please wait)')
        hippo.bet(horseNumber, {
            gas: 300000,
            from: account,
            value: web3.toWei(1, 'ether')
        })
            .then(function (result) {
                self.updateAll(result)
            })
            .catch(function (e) {
                console.log(e)
                self.updateStatus('Error sending transaction; see log.')
            })
    },

    updateAll: function (result) {
        for (var i = 0; i < result.logs.length; i++) {
            var log = result.logs[i]
            if (log.event == 'HippoEvent') {
                // We found the event!
                console.log(log)
                var minimumBet = log.args._minimumBet.c[0]
                var totalBet = log.args._totalBet.c[0]
                var numberOfBets = log.args._numberOfBets.c[0]
                var maxPalayers = log.args._maxPlayers.c[0]
                var raceId = log.args._raceId.c[0]
                console.log(raceId, minimumBet, totalBet, numberOfBets, maxPalayers)
                document.getElementById('minimumBet').innerHTML = minimumBet.valueOf()
                document.getElementById('totalBet').innerHTML = totalBet.valueOf()
                document.getElementById('numberOfBets').innerHTML = numberOfBets.valueOf()
                document.getElementById('maxPlayers').innerHTML = maxPalayers.valueOf()
                document.getElementById('raceId').innerHTML = raceId.valueOf()
                // this.updateCurrentAcc()
                break
            }
        }
    }

}

window.addEventListener('load', function () {
    // Checking if Web3 has been injected by the browser (Mist/MetaMask)
    if (typeof web3 !== 'undefined') {
        console.warn("Using web3 detected from external source. If you find that your accounts don't appear or you have 0 MetaCoin, ensure you've configured that source properly. If using MetaMask, see the following link. Feel free to delete this warning. :) http://truffleframework.com/tutorials/truffle-and-metamask")
        // Use Mist/MetaMask's provider
        window.web3 = new Web3(web3.currentProvider)
    } else {
        console.warn("No web3 detected. Falling back to http://127.0.0.1:9545. You should remove this fallback when you deploy live, as it's inherently insecure. Consider switching to Metamask for development. More info here: http://truffleframework.com/tutorials/truffle-and-metamask")
        // fallback - use your fallback strategy (local node / hosted node + in-dapp id mgmt / fail)
        window.web3 = new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:9545'))
    }

    App.start()
})
© 2018 GitHub, Inc.
    Terms
Privacy
Security
Status
Help
Contact GitHub
API
Training
Shop
Blog
About