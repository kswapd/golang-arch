console.log("Hello from my.js");


function formatDate(date, format) {
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
async function getData() {
    document.getElementById('data').textContent = "Loading...";
    const response = await fetch('/api/data');
    //const rawData = await response.text();
    //console.log(rawData)
    const data = await response.json();
    console.log(data)
    if (data["messageType"] == "1000") {
        const milliseconds = data["body"]["tradeTime"]; // Example milliseconds value
        const date = new Date(milliseconds);
        //console.log(date.toString());
        const formattedDate = formatDate(date, 'yyyy-mm-dd HH:MM:SS');
        document.getElementById('data').innerHTML = "名称: " + data["body"]["name"] + "<br>";
        document.getElementById('data').innerHTML += "交易时间: " + formattedDate + "<br>";
        document.getElementById('data').innerHTML += "价格:" + data["body"]["price"] + "(" + data["body"]["unit"] + ")";
    }
}
document.addEventListener('DOMContentLoaded', function () {
    // Your code to run after the DOM is ready
    getData();
});
