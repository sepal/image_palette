import {Component} from '@angular/core';

@Component({
    styles: [require('./upload.component.scss')],
    template: require('./upload.component.html')
})
export class UploadComponent {
    image: File;
    image_submitted: boolean;
    image_src: string;
    image_name: string;

    colors: Array<string>;


    constructor() {
        this.image_submitted = false;
    }

    uploadImage() {
        this.makeFileRequest('http://localhost:8000/upload', [], this.image).then((result: any) => {
            this.image_submitted = true;

            let ws: WebSocket = new WebSocket(`ws://localhost:8000/image/${result.id}`);
            ws.onmessage = (event) => {
                let img = JSON.parse(event.data);
                this.colors = img.colors;
            };
        }).catch((err) => {
            console.error('Error while uploading image:');
            console.error(err);
        });
    }

    fileSelected(event) {
        let file: File = event.target.files[0];
        if (event.target.files.length < 1) {
            return;
        }
        if (file.type !== 'image/jpeg' && file.type !== 'image/png') {
            alert('Wrong file!');
            this.image_src = null;
            return;
        }
        let reader: FileReader = new FileReader();

        reader.addEventListener('load', (e: any) => {
            this.image_src = e.target.result;
            this.image_name = this.image.name;
        }, false);
        this.image = file;

        reader.readAsDataURL(this.image);

    }

    makeFileRequest(url: string, params: Array<string>, file: File) {
        return new Promise((resolve, reject) => {
            let formData: any = new FormData();
            let xhr = new XMLHttpRequest();
            formData.append('image', file, file.name);

            xhr.onreadystatechange = function () {
                if (xhr.readyState === 4) {
                    if (xhr.status === 200) {
                        resolve(JSON.parse(xhr.response));
                    } else {
                        reject(xhr.response);
                    }
                }
            };

            xhr.open('POST', url, true);
            xhr.send(formData);
        });
    }
}
