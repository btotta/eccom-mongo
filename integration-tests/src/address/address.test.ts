import request from 'supertest';
import { AddressRequest } from './model/address_request';

declare const BASE_URL: string;
declare const TEST_EMAIL: string;
declare const TEST_PASSWORD: string;

describe('Testing create address', () => {
    it('Should return status 200 and a JSON object', async () => {

        const req: AddressRequest = {
            city: "sao paulo",
            complement: "complement",
            country: "brasil",
            neighborhood: "parque munhoz",
            number: "11",
            state: "sao paulo",
            street: "rua 1",
            zip_code: "12345678"
        };

        const logres = await request(BASE_URL).post('/user/login').send({ email: TEST_EMAIL, password: TEST_PASSWORD });
        expect(logres.status).toBe(200);

        const res = await request(BASE_URL).post('/address').send(req).set('Authorization', `${logres.body.token}`);

        expect(res.status).toBe(200);

    });
});