

export default [
    {
        url: '/status',
        method: 'get',
        response: () => {
            return {
                "result": true,
                "data": {
                    "version": "0.0.1",
                    "start_time": "2024-01-01"
                }
            }
        },
    },
];
