import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllCommentsAdminComponent } from './all-comments-admin.component';

describe('AllCommentsAdminComponent', () => {
  let component: AllCommentsAdminComponent;
  let fixture: ComponentFixture<AllCommentsAdminComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AllCommentsAdminComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AllCommentsAdminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
