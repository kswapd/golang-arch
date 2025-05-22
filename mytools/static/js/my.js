console.log("Hello from my.js");

async function getData() {
    document.getElementById('data').textContent = "Loading...";
    const response = await fetch('/api/data');
    //const rawData = await response.text();
    //console.log(rawData)
    const data = await response.json();
    console.log(data)
    document.getElementById('data').textContent = "名称: " + data["body"]["name"] + "";
    document.getElementById('data').textContent += "交易时间: " + data["body"]["tradeTime"] + "";
    document.getElementById('data').textContent += "价格:" + data["body"]["price"] + "(" + data["body"]["unit"] + ")";
}
document.addEventListener('DOMContentLoaded', function () {
    // Your code to run after the DOM is ready
    getData();
});
