
export async function getBackendStatus() {
    const res = await fetch(`/status`);
    return res.json();
}