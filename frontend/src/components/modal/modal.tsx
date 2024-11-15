import { IModal } from '$types/modal'

import '@styles/components/modal/modal.css'

export default function Modal({ active, setActive, children, width }: IModal) {
    return (
        <div className={active ? "modal flex center active" : "modal flex center"} onClick={() => setActive(false)}>
            <div style={{ width: `${width}px` }} className={active ? "modal__content bgc-2 b active" : "modal__content b"} onClick={e => e.stopPropagation()}>
                { children }
            </div>
        </div>
    )
}