import { TestBed } from '@angular/core/testing';
import { BookService } from './book.service';
describe('BookService', () => {
    beforeEach(() => TestBed.configureTestingModule({}));
    it('should be created', () => {
        const service = TestBed.get(BookService);
        expect(service).toBeTruthy();
    });
});
//# sourceMappingURL=book.service.spec.js.map