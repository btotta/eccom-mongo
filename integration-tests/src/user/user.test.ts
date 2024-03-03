import request from 'supertest';
import { CreateUserRequest } from './model/user_request';

declare const BASE_URL: string;
declare const TEST_EMAIL: string;
declare const TEST_PASSWORD: string;

describe('Testing the creation of a user', () => {

    it('With valid credentials, should return status 200 and a JSON object', async () => {

        const user: CreateUserRequest = {
            confirm_password: "E!ZGiNA29$zmnqH$",
            document: Math.floor(10000000 + Math.random() * 90000000).toString(),
            email: "Test" + Math.floor(10000000 + Math.random() * 90000000).toString() + "@gmail.com",
            password: "E!ZGiNA29$zmnqH$",
            last_name: "Automated",
            name: "Test"
        };

        const res = await request(BASE_URL).post('/user').send(user);

        expect(res.status).toBe(200);
        expect(res.body).toHaveProperty("id");
        expect(res.body).toHaveProperty("email");
        expect(res.body).toHaveProperty("name");
        expect(res.body).toHaveProperty("last_name");
    });

    it('With invalid email, should return status 400 and a JSON object', async () => {

        const user: CreateUserRequest = {
            confirm_password: "E!ZGiNA29$zmnqH$",
            document: Math.floor(10000000 + Math.random() * 90000000).toString(),
            email: "Test" + Math.floor(10000000 + Math.random() * 90000000).toString() + "gmail.com",
            password: "E!ZGiNA29$zmnqH$",
            last_name: "Automated",
            name: "Test"
        };

        const res = await request(BASE_URL).post('/user').send(user);

        expect(res.status).toBe(400);
        expect(res.body).toHaveProperty("error", "this email cannot be used");
    });

    it('With invalid password, should return status 400 and a JSON object', async () => {

        const user: CreateUserRequest = {
            confirm_password: "123321131",
            document: Math.floor(10000000 + Math.random() * 90000000).toString(),
            email: "Test" + Math.floor(10000000 + Math.random() * 90000000).toString() + "@gmail.com",
            password: "asfcasf123",
            last_name: "Automated",
            name: "Test"
        };

        const res = await request(BASE_URL).post('/user').send(user);

        expect(res.status).toBe(400);
        expect(res.body).toHaveProperty("error", "this password cannot be used");
    });

});


describe('Testing the login of a user', () => {

    it('With valid credentials, should return status 200 and a JSON object', async () => {

        const user: CreateUserRequest = {
            confirm_password: "E!ZGiNA29$zmnqH$",
            document: Math.floor(10000000 + Math.random() * 90000000).toString(),
            email: "automated_test@gmail.com",
            password: "E!ZGiNA29$zmnqH$",
            last_name: "Automated",
            name: "Test"
        };

        await request(BASE_URL).post('/user').send(user);

        const res = await request(BASE_URL).post('/user/login').send({ email: TEST_EMAIL, password: TEST_PASSWORD });

        expect(res.status).toBe(200);
        expect(res.body).toHaveProperty("refresh_token");
        expect(res.body).toHaveProperty("token");
    });

    it('With invalid email, should return status 400 and a JSON object', async () => {

        const res = await request(BASE_URL).post('/user/login').send({ email: "sadfasfaf@gmail.com", password: "asfasf" });

        expect(res.status).toBe(400);
        expect(res.body).toHaveProperty("error", "invalid login credentials");

    });

});
