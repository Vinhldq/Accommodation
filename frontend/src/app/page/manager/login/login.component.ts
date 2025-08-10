import { Component, OnInit } from '@angular/core';
import {
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    Validators,
} from '@angular/forms';
import { TuiIcon, TuiTextfield } from '@taiga-ui/core';
import { TuiPassword } from '@taiga-ui/kit';
import { AuthService } from '../../../services/manager/auth.service';
import { ManagerLoginInput } from '../../../models/manager/accommodation.model';
import { Router } from '@angular/router';
import { SaveTokenToCookie } from '../../../shared/token/token';
import { MessageService } from 'primeng/api';
import { Toast } from 'primeng/toast';
import { ButtonModule } from 'primeng/button';
import { LoaderComponent } from '../../../components/loader/loader.component';
import { finalize } from 'rxjs';

@Component({
    selector: 'app-login',
    imports: [
        TuiTextfield,
        TuiIcon,
        ReactiveFormsModule,
        TuiPassword,
        Toast,
        ButtonModule,
        LoaderComponent,
    ],
    templateUrl: './login.component.html',
    styleUrl: './login.component.scss',
    providers: [MessageService],
})
export class LoginComponent implements OnInit {
    protected formLogin = new FormGroup({
        account: new FormControl('', Validators.required),
        password: new FormControl('', Validators.required),
    });
    isLoading: boolean = false;

    constructor(
        private authSerivce: AuthService,
        private router: Router,
        private messageService: MessageService
    ) {}
    ngOnInit(): void {
        this.formLogin.get('account')?.valueChanges.subscribe(() => {
            if (this.formLogin.get('account')?.hasError('backend')) {
                this.formLogin.get('account')?.setErrors(null);
                this.formLogin.get('account')?.updateValueAndValidity();
                this.formLogin.get('password')?.updateValueAndValidity();
            }
        });

        this.formLogin.get('password')?.valueChanges.subscribe(() => {
            if (this.formLogin.get('account')?.hasError('backend')) {
                this.formLogin.get('account')?.setErrors(null);
                this.formLogin.get('account')?.updateValueAndValidity();
                this.formLogin.get('password')?.updateValueAndValidity();
            }
        });
    }

    showToast(
        severity: 'success' | 'info' | 'warn' | 'error',
        summary: string,
        detail: string
    ): void {
        this.messageService.add({ severity, summary, detail });
    }

    handleLogin() {
        if (this.formLogin.invalid) {
            this.formLogin.markAllAsTouched();
            return;
        }
        this.isLoading = true;

        let managerLogin: ManagerLoginInput = {
            account: this.formLogin.value.account ?? '',
            password: this.formLogin.value.password ?? '',
        };

        this.authSerivce
            .login(managerLogin)
            .pipe(finalize(() => (this.isLoading = false)))
            .subscribe({
                next: (response) => {
                    SaveTokenToCookie(response.data.token);
                    localStorage.setItem(
                        'managerUserName',
                        response.data.user_name
                    );

                    this.router.navigate(['/manager/accommodation']);
                },
                error: (error) => {
                    const errorMessage =
                        error.error.message ||
                        'Đã xảy ra lỗi trong quá trình đăng nhập. Vui lòng thử lại sau.';
                    this.formLogin
                        .get('account')
                        ?.setErrors({ backend: errorMessage });
                },
            });
    }
}
