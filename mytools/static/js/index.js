

var app = new Vue({
    el: '.main-container',
    data: {
        message: 'Hello Vue!',
        menuItems: [
            {
                text: 'Home',
                isOpen: false,
                children: []
            },
            {
                text: 'About',
                isOpen: false,
                children: [
                    { text: 'Team', isOpen: false, children: [] },
                    { text: 'Company', isOpen: false, children: [] }
                ]
            },
            {
                text: 'Services',
                isOpen: false,
                children: [
                    { text: 'Web Development', isOpen: false, children: [] },
                    { text: 'Mobile Development', isOpen: false, children: [] }
                ]
            },
            {
                text: 'Contact',
                isOpen: false,
                children: []
            }
        ],
        currentView: 'Home' // Default view
    },
    mounted: function () {
        console.log('dddd')
        this.getData();
    },
    methods: {
        switchView: function (view) {
            this.currentView = view;
        },
        toggleMenu: function (item) {
            item.isOpen = !item.isOpen;
        },
        handleMenuClick(item) {
            if (item.children && item.children.length > 0) {
                // If the item has children, toggle its `isOpen` property
                item.isOpen = !item.isOpen;
            } else {
                // If the item has no children, switch the view
                this.currentView = item.text;
            }
        },
        formatDate: function (date, format) {
            const map = {
                'mm': ('0' + (date.getMonth() + 1)).slice(-2),
                'dd': ('0' + date.getDate()).slice(-2),
                'yyyy': date.getFullYear(),
                'HH': ('0' + date.getHours()).slice(-2),
                'MM': ('0' + date.getMinutes()).slice(-2),
                'SS': ('0' + date.getSeconds()).slice(-2),
            };
            return format.replace(/mm|dd|yyyy|HH|MM|SS/gi, matched => map[matched]);
        },
        getData: async function () {
            this.message = "Loading...";
            const response = await fetch('/api/data');
            const data = await response.json();
            console.log(data)
            if (data["messageType"] == "1000") {
                const milliseconds = data["body"]["tradeTime"];
                const date = new Date(milliseconds);
                const formattedDate = this.formatDate(date, 'yyyy-mm-dd HH:MM:SS');
                this.message = "名称: " + data["body"]["name"] + "<br>";
                this.message += "交易时间: " + formattedDate + "<br>";
                this.message += "价格:" + data["body"]["price"] + "(" + data["body"]["unit"] + ")";
            }
        }
    }
});



