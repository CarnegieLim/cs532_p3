import request from '@/utils/request';
export async function predict(sentence) {
  return request('/predict?sentence' + sentence);
}
