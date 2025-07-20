const Home = {
    template: ` 
        <div>
            <h1>Welcome to Home</h1>
            <p>This is the Home page content.</p>
            <div class="gold-info">
                <button @click="getData">Refresh Data</button>
                <div v-html="message"></div>
            </div>
        </div>
    `,
    data() {
        return {
            message: 'Hello Vue!'
        }
    },
    mounted: function () {
        console.log('eee')
        //this.handleInitialView();
        this.getData();
    },
    methods: {
        async getData() {
            this.message = "Loading...";
            const response = await fetch('/api/data');
            const data = await response.json();
            if (data["messageType"] == "1000") {
                const milliseconds = data["body"]["tradeTime"];
                const date = new Date(milliseconds);
                const formattedDate = this.formatDate(date, 'yyyy-mm-dd HH:MM:SS');
                this.message = "名称: " + data["body"]["name"] + "<br>";
                this.message += "交易时间: " + formattedDate + "<br>";
                this.message += "价格:" + data["body"]["price"] + "(" + data["body"]["unit"] + ")";
            }
        },
        formatDate(date, format) {
            const map = {
                'mm': ('0' + (date.getMonth() + 1)).slice(-2),
                'dd': ('0' + date.getDate()).slice(-2),
                'yyyy': date.getFullYear(),
                'HH': ('0' + date.getHours()).slice(-2),
                'MM': ('0' + date.getMinutes()).slice(-2),
                'SS': ('0' + date.getSeconds()).slice(-2),
            };
            return format.replace(/mm|dd|yyyy|HH|MM|SS/gi, matched => map[matched]);
        }
    }
}