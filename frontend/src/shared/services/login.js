import axios from 'axios';

import { SERVER_API_URL } from '../../config/constants';

export const CheckLoginToken = async userToken => 
    await axios.post(`${SERVER_API_URL}/login`, userToken)