import {Component} from '@angular/core';

@Component({
    styles: [`
    img {
        max-width: 960px;
        height: auto;
    }
    .color {
        display: inline-block;
        width: 100px;
        height: 100px;
        margin: 5px;
        text-shadow: 0px 0px 5px rgba(0, 0, 0, 0.8);
        
    }
    
    .color span {
        height: 100px;
        line-height: 100px;
    }
    `],
    template: `
    <div *ngIf="image_submitted == false">
        <input (change)="fileSelected($event)" type="file"/>
        <input (click)="uploadImage($event)" type="submit" />
    </div>
    <div *ngIf="image_submitted == true">
        <div *ngFor="let color of colors" class="color" [style.background-color]="color">
          <span>{{color}}</span>
        </div>
    </div>
    <div *ngIf="image_src != null">
        <img class="preview" src={{image_src}} alt={{image_name}}>
    </div>
    `
})
export class UploadComponent {
    image: File;
    image_submitted: boolean;
    image_src: string;
    image_name: string;

    colors: Array<string>;


    constructor() {
        this.image_submitted = false
    }

    uploadImage() {
        this.makeFileRequest('http://localhost:8000/upload', [], this.image).then((result) => {
            this.server_response = result.id;
            this.image_submitted = true;

            let ws: WebSocket = new WebSocket(`ws://localhost:8000/image/${result.id}`);
            ws.onmessage = (event) => {
                let img = JSON.parse(event.data);
                this.colors = img.colors;
            }
        }).catch((err) => {
            console.error("Error while uploading image:");
            console.error(err);
        });
    }

    fileSelected(event) {
        let file: File = event.target.files[0];
        if (event.target.files.length < 1) {
            return
        }
        if (file.type !== 'image/jpeg' && file.type !== 'image/png') {
            alert('Wrong file!');
            this.image_src = null;
            return
        }
        let reader: FileReader = new FileReader();

        reader.addEventListener("load", (event) => {
            this.image_src = event.target.result;
            this.image_name = this.image.name;
        }, false);
        this.image = file;

        reader.readAsDataURL(this.image);

    }

    makeFileRequest(url: string, params: Array<string>, file: File) {
        return new Promise((resolve, reject) => {
            let formData: any = new FormData();
            let xhr = new XMLHttpRequest();
            formData.append("image", file, file.name);

            xhr.onreadystatechange = function () {
                if (xhr.readyState == 4) {
                    if (xhr.status == 200) {
                        resolve(JSON.parse(xhr.response));
                    } else {
                        reject(xhr.response);
                    }
                }
            };

            xhr.open("POST", url, true);
            xhr.send(formData);
        });
    }
}
