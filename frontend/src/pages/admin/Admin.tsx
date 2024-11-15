import Header from "@components/header/header.tsx";
import Loading from "../../lazyLoad.tsx"
import { useEffect, useState } from "react";
import { IUser } from "@api/axios/entities/authResponse.ts";
import AdminService from "@api/axios/requests/admin.ts";

import '@styles/pages/admin/admin.css'

export default function Admin() {
    const [loading, setLoading] = useState(true);
    const [users, setUsers] = useState<IUser[]>([]);

    useEffect(() => {
        async function getUsers() {
            const response = await AdminService.getUsers()
            setUsers(response.data.data)
            setLoading(false)
        }
        getUsers()
    }, [])

    const deleteUser = async (id: string) => {
        const response = await AdminService.deleteUser(id)
        console.log(response)
        window.location.reload();
    }

    const updateRole = async (id: string, role: string) => {
        const response = await AdminService.updateRole(id, role)
        console.log(response)

    }

    return (
        <div className="container">
            <Header />
            {loading ? <Loading/> : <div className='admin-container w100'>
                <div className='admin-box b bgc-1 br10 w100 h100'>
                    <div className='admin-inner w100 h100'>
                        <div className={'admin-categories flex'}>
                            <div className='admin-category flex center b br10 bgc-2'>Никнейм</div>
                            <div className='admin-category flex center b br10 bgc-2'>Почта</div>
                            <div className='admin-category flex center b br10 bgc-2'>Айпи</div>
                            <div className='admin-category flex center b br10 bgc-2'>Роль</div>
                        </div>
                        <div className='user-items flex col'>
                            {
                                users.map(({user_id, nickname, email, ip_addr, role}, index) => {
                                    return (
                                        <div className={'user-item flex t-center'} key={index}>
                                            <p className={'item item-nickname flex center'}>{nickname}</p>
                                            <p className={'item item-email flex center'}>{email}</p>
                                            <p className={'item item-ipaddr flex center'}>{ip_addr}</p>
                                            <form className={'item item-form'}>
                                                <select className={'item item-roles bgc-2 b br10'} defaultValue={role} onChange={(e) => updateRole(user_id, e.target.value)}>
                                                    <option className={'item-roles-option'} value="player">Игрок</option>
                                                    <option className={'item-roles-option'} value="admin">Админ</option>
                                                </select>
                                            </form>
                                            <p className={'delete-user b bgc-2 br10'} onClick={() => deleteUser(user_id)}></p>
                                        </div>
                                    )
                                })
                            }
                        </div>
                    </div>
                </div>
            </div>}
        </div>
    )
}