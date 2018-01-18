/**
 * Created by zim on 17-Jan-18.
 */
new Vue ({
        el: '#app',
        // mounted: function () {
        //     this.fetchAssets();
        // },
        data: {
            address: '',
            sellerAddress: '',
            message: 'Click for new payment address',
            balance: -1,
            statusMessage: '',
            checked: false
        },
        methods:{
            checkStatus: function () {
                this.statusMessage = "pending..."
                this.checked = true;
                console.log(this.checked)
                fetch("/status")
                    .then(response => response.json())
                .then(json => {
                    console.log(json);
                this.address = json.address;
                this.statusPaid = json.status;
                this.balance = json.balance;
                if (this.statusPaid){
                    this.statusMessage = "Your payment was made"}
                else {this.statusMessage = "Your payment wasn't found"}
            })
            },

            getAddress: function () {
                fetch("/address")
                    .then(response => response.json())
                .then(json => {
                    console.log(json);
                this.address = json.address;
            })
            }

        }


    }
);