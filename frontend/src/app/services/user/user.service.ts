import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { UpdateUser, UserResponse } from '../../models/user/user.model';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';

@Injectable({
    providedIn: 'root',
})
export class UserService {
    private apiUrl = `${environment.apiUrl}/user`;

    constructor(private http: HttpClient) {}

    getUserInfo(): Observable<UserResponse> {
        return this.http.get<UserResponse>(`${this.apiUrl}/get-user-info`);
    }

    updateUserInfo(userData: UpdateUser): Observable<UserResponse> {
        const newUser: UpdateUser = {
            username: userData.username,
            phone: userData.phone,
            gender: userData.gender,
            birthday: userData.birthday,
        };
        return this.http.post<UserResponse>(
            `${this.apiUrl}/update-user-info`,
            newUser
        );
    }
}
