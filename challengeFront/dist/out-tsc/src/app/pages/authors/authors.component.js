import * as tslib_1 from "tslib";
import { Component } from '@angular/core';
let AuthorsComponent = class AuthorsComponent {
    constructor(authorService) {
        this.authorService = authorService;
        this.authors = [];
    }
    ngOnInit() {
        this.authorService.getAllAuthors().subscribe((response) => {
            this.authors = response;
        });
    }
};
AuthorsComponent = tslib_1.__decorate([
    Component({
        selector: 'app-authors',
        templateUrl: './authors.component.html',
        styleUrls: ['./authors.component.scss']
    })
], AuthorsComponent);
export { AuthorsComponent };
//# sourceMappingURL=authors.component.js.map