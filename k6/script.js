import http from 'k6/http';
import { sleep } from 'k6';
export let options = {
    vus: 10,
    duration: '10s',
};
export default function () {
    http.get('http://localhost:8081/get?from=2019-02-01&to=2021-02-04&space=Gomel_BrL2_turnstile');
    sleep(1);
}
//