# Repo SDM
## Description
This repository contains the code for the paper "A Study of the Impact of Data Augmentation on the Performance of Deep Learning Models for the Detection of COVID-19 in Chest X-ray Images". The paper is currently under review. The code is written in Python and uses the PyTorch library. The code is divided into two main parts: the data augmentation techniques and the deep learning models. The data augmentation techniques are implemented in the `data_augmentation` folder and the deep learning models are implemented in the `models` folder. The code is organized in the following way:
- `data_augmentation` folder: contains the code for the data augmentation techniques.
- `models` folder: contains the code for the deep learning models.
- `utils` folder: contains utility functions used in the code.
  - `config.py`: contains the configuration parameters used in the code.
  - `data.py`: contains the data loading functions.
  - `evaluate.py`: contains the evaluation functions.
  - `train.py`: contains the training functions.

## Data Augmentation Techniques
The data augmentation techniques implemented in this repository are:
- Rotation
- Translation
- Shear
- Zoom
- Horizontal Flip
- Vertical Flip
- Brightness
- Contrast
- Saturation
- Hue
- Gaussian Noise

## Deep Learning Models
The deep learning models implemented in this repository are:
- AlexNet
- VGG16
- ResNet18
- DenseNet121

## Usage
To use the code, follow these steps:
1. Clone the repository:
```bash
git clone
``` 
2. Install the required packages:
```bash
pip install -r requirements.txt
```
3. Run the code:
```bash
python main.py
```

## Results
The results of the experiments are stored in the `results` folder. The results are stored in CSV files with the following columns:
- Model: the name of the deep learning model.
- Data Augmentation: the name of the data augmentation technique.
- Accuracy: the accuracy of the model on the test set.
- Precision: the precision of the model on the test set.
- Recall: the recall of the model on the test set.
- F1 Score: the F1 score of the model on the test set.
- AUC: the area under the ROC curve of the model on the test set.
- Training Time: the training time of the model in seconds. 

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


