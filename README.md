# Image editing website
![go-version](https://img.shields.io/github/go-mod/go-version/NooFreeNames/ImageEditor?style=flat-square)

The website is a convenient tool for image editing. With the help of this site, users can crop and filter images in JPG and PNG formats. A feature of this site is the use of goroutin for image processing, which reduces the waiting time and speeds up the editing process. Thanks to the simple and intuitive interface, users can quickly and easily edit their images, while getting high-quality results.

## Technologies Used

- html
- css
- js
- bootstrap
- golang
- godotenv
- testify

## Configuring the Server

The server is configured using the ./configs/.env file.

File contents:
```env
SERVER_HOST=127.0.0.1
SERVER_PORT=8080
SITE_DIR=web
```

* SERVER_HOST – the address of the host on which the server will be launched.
* SERVER_PORT – the port on which the server will be started.
* SITE_DIR – the directory where the site files are located.

## Installation

1. Clone the repository: `git clone https://github.com/NooFreeNames/ImageEditor.git`.
2. Go to the project directory.
3. Installing the necessary modules: `go mod download`.
4. Start the server: `go run ./cmd/main.go`.

## Usage

### How to open a website?

After starting the server, click on the link that is displayed in the console.

![website link](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/0.png)<br>

The main page should open.

![website](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/1.png)

### Changing the site theme

Click on the green button on the top right to change the theme.

![theme button](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/2.png)<br>
![light theme](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/3.png)

### Uploading an image

To upload an image, click on the file selection field and upload the desired image.

![uploading image](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/4.png)<br>
![image uploaded](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/5.png)

### Cropping an image

To crop the image, specify the new crop size and position. Then click Submit.

![cropping image](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/6.png)<br>
![image cropped](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/7.png)

### Image Filtering

To apply a filter to an image, select it from the drop-down list and click the Submit.

![image filtering](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/8.png)<br>
![image filtered](https://github.com/NooFreeNames/ImageEditor/blob/master/readme-images/9.png)
