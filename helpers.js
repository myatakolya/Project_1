export default function debug(obj = {}) {
  return JSON.stringify(obj, null, 4)
}