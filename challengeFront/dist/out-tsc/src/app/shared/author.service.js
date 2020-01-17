import * as tslib_1 from "tslib";
import { Injectable } from '@angular/core';
import { tap } from 'rxjs/operators';
import { map } from 'rxjs/operators';
let AuthorService = class AuthorService {
    constructor(http) {
        this.http = http;
    }
    getPosts() {
        return this.http.get('localhost:8000')
            .pipe(tap(console.log));
    }
    updateAuthor(multipart, headers) {
        return this.http.patch('localhost:8000/authors', multipart, headers)
            .pipe(map((data) => data));
    }
    createAuthor(multipart, headers) {
        return this.http.post('localhost:8000/authors', multipart, headers)
            .pipe(map((data) => data));
    }
    getAllAuthors(headers) {
        return this.http.get('http://localhost:8000/authors', headers)
            .pipe(map((data) => data));
    }
};
AuthorService = tslib_1.__decorate([
    Injectable({
        providedIn: 'root'
    })
], AuthorService);
export { AuthorService };
//# sourceMappingURL=author.service.js.map