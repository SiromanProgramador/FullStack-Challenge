import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { tap } from 'rxjs/operators';
import { map } from 'rxjs/operators';


@Injectable({
  providedIn: 'root'
})
export class AuthorService {

  constructor( private http: HttpClient ) { }


    getPosts() {
        return this.http.get('localhost:8000')
            .pipe( tap( console.log ) );
    }

    updateAuthor(multipart: any, headers?: any) {
        return this.http.patch<any>('http://localhost:8000/authors', multipart, headers)
        .pipe(
            map((data: any) => data)
        );
    }
    createAuthor(multipart: any, headers?: any) {
        return this.http.post<any>('http://localhost:8000/authors', multipart, headers)
        .pipe(
            map((data: any) => data)
        );
    }
    getAllAuthors(headers?: any) {
        return this.http.get<any>('http://localhost:8000/authors',headers)
        .pipe(
            map((data: any) => data)
        );

    }
}