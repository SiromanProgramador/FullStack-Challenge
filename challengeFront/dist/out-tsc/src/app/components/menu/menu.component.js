import * as tslib_1 from "tslib";
import { Component } from '@angular/core';
let MenuComponent = class MenuComponent {
    constructor() {
        this.rutas = [
            {
                name: 'Books',
                path: '/books'
            },
            {
                name: 'Authors',
                path: '/authors'
            },
        ];
    }
    ngOnInit() {
    }
};
MenuComponent = tslib_1.__decorate([
    Component({
        selector: 'app-menu',
        templateUrl: './menu.component.html',
        styleUrls: ['./menu.component.scss']
    })
], MenuComponent);
export { MenuComponent };
//# sourceMappingURL=menu.component.js.map