import axios from 'axios';

import { SERVER_API_URL } from '../../config/constants';

export const CheckLoginAccountId = async userAccountId => 
    await axios.post(`${SERVER_API_URL}/login`, userAccountId)

export const GetUserDetails = async userAccountId => 
    await axios.post(`${SERVER_API_URL}/user`, userAccountId)