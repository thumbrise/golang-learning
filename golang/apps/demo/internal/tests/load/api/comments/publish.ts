import http from 'k6/http';
import {Options} from 'k6/options';
import exec from 'k6/execution';

export const options: Options = {
    summaryTrendStats: ['avg', 'med', 'p(90)', 'p(95)', 'p(99)'],
    scenarios: {
        'base': {
            executor: "constant-arrival-rate",
            rate: 3000,
            duration: '30s',
            preAllocatedVUs: 10,
        }
    }

}

interface Payload {
    userUUID: string
    postUUID: string
    content: string
}

function payload() {
    let id = exec.vu.iterationInInstance
    return {
        userUUID: `user${id}`,
        postUUID: `post${id}`,
        content: `contentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontentcontent${id}`
    }
}

export default function () {
    let body = JSON.stringify(payload());

    let headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    }
    http.post(`http://app:8080/api/comments/publish`, body, {
        headers: headers
    })
}
