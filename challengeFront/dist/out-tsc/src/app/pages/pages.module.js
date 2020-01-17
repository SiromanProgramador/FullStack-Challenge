import * as tslib_1 from "tslib";
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AuthorsComponent } from './authors/authors.component';
import { BooksComponent } from './books/books.component';
let PagesModule = class PagesModule {
};
PagesModule = tslib_1.__decorate([
    NgModule({
        declarations: [
            AuthorsComponent,
            BooksComponent
        ],
        exports: [
            AuthorsComponent,
            BooksComponent
        ],
        imports: [
            CommonModule
        ]
    })
], PagesModule);
export { PagesModule };
//# sourceMappingURL=pages.module.js.map