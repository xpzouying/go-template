import { getBackendStatus } from './api';

export async function loader() {
    const backendStatus = await getBackendStatus();
    return backendStatus;
}
