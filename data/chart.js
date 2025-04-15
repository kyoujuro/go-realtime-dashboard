const ctx = document.getElementById('chart').getContext('2d');
const labels = [];
const dataPoints = [];

const chart = new Chart(ctx, {
    type: 'line',
    data: {
        labels: labels,
        datasets: [{
            label: 'Live Value',
            data: dataPoints,
            borderWidth: 2,
            borderColor: 'blue',
            fill: false
        }]
    },
    options: {
        animation: false,
        scales: {
            x: { title: { display: true, text: 'Time' }},
            y: { beginAtZero: true }
        }
    }
});

const ws = new WebSocket("ws://" + location.host + "/ws");

ws.onmessage = (event) => {
    const data = JSON.parse(event.data);
    labels.push(data.timestamp);
    dataPoints.push(data.value);

    if (labels.length > 20) { // 最新20件だけ保持
        labels.shift();
        dataPoints.shift();
    }

    chart.update();
};
