from flask import Flask, request, jsonify
import cv2
import numpy as np
import tensorflow as tf  # Ganti ini sesuai framework model Anda (PyTorch, Scikit-learn, dll.)

app = Flask(__name__)

model = tf.keras.models.load_model("waste_model/model/waste.keras")
secretKey = "dlahsldhlahdq29o3u01jhasdlajdkajsdkljasdkjahsdklahfklashf"

def preprocess_image(image_path):
    img_array = cv2.imread(image_path)
    img_array = cv2.cvtColor(img_array, cv2.COLOR_BGR2RGB)
    img_array = cv2.resize(img_array, (224, 224))
    img_array = img_array / 255.0
    img_array = np.expand_dims(img_array, axis=0) 
    return img_array

@app.route('/predict', methods=['POST'])
def predict():
    try:
        data = request.get_json()
        headers = request.headers

        if headers["AUTH"] != secretKey:
            return jsonify({
                "status": "Unatuhorized",
            }), 401

        path = data["path_img"]
        processed_image = preprocess_image(path)
        
        predictions = model.predict(processed_image)
        predicted_class = np.argmax(predictions)  

        waste = ""
        if predicted_class == 0:
            waste = "organic"
        else:
            waste = "non-organic"
        
        return jsonify({
            "status": "success",
            "prediction": waste
        }), 200
    except Exception as e:
        return jsonify({
            "status": "error",
            "message": str(e)
        }), 500

if __name__ == '__main__':
    app.run(host='localhost', port=8081)
