from flask import Flask, jsonify, request

app = Flask(__name__)

location_data = {
    "latitude": None,
    "longitude": None
}

@app.route("/")
def index():
    return """
    <h2>Get Classroom GPS</h2>
    <button onclick="getLocation()">Send GPS</button>
    <script>
    function getLocation() {
    navigator.geolocation.getCurrentPosition(function(position) {
    fetch(window.location.origin + '/update', {
    method: 'POST',
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify({
    latitude: position.coords.latitude,
    longitude: position.coords.longitude
    })
    }).then(res => res.json())
    .then(data => console.log(data));
    alert("GPS sent to microservice!");
    });
    }
    </script>
    """

@app.route("/update", methods=["POST"])
def update():
    data = request.json
    location_data["latitude"] = data["latitude"]
    location_data["longitude"] = data["longitude"]
    return {"status": "updated"}

@app.route("/location")
def location():
    return jsonify(location_data)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)

