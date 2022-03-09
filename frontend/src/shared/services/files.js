import axios from 'axios';

import { SERVER_API_URL } from '../../config/constants';

export const GetFilesMigratedForUser = async accountId => 
    await axios.post(`${SERVER_API_URL}/files`, accountId)