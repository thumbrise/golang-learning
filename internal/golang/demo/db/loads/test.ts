import http from 'k6/http';
import {Options} from 'k6/options';

export const options:Options = {
    vus: 10,
    duration: '10s',
    summaryTrendStats: ['avg', 'med', 'p(90)', 'p(95)', 'p(99)'],
}
export default function () {
    let payload = {
        email: "murat@k6.......io",
    };
    let headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    }
    http.post(`http://app:8080/api/auth/sign-in`, JSON.stringify(payload), {
        headers: headers
    })
}
