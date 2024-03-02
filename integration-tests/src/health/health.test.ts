import request from 'supertest';

declare const BASE_URL: string;

describe('Testing the health api', () => {
    it('Should return status 200 and a JSON object', async () => {
        const res = await request(BASE_URL).get('/health');
        expect(res.status).toBe(200);

        expect(res.body).toHaveProperty("message", "It\'s healthy");
    });
});