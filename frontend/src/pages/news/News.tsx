import Header from "../../components/header/header.jsx";

import './news.css'

export default function News() {
    return (
        <>
            <Header />
            <div className={'new-container'}>
                <div className={'news-inner-container flex w100 h100'}>
                    <aside className={'b br20 bgc-1 news-filters h100'}>
                        <div className={'news-filters__list'}>
                            <div>
                                <a className={'news-filters-link'} href="">
                                    <i className={'news-filters__icon'}></i>
                                    <span className={'news-filters-text'}>Какой-то фильтр</span>
                                </a>
                                <div className={'news-filters__content'}>
                                    <div className={'news-filters-checkbox-group'}>
                                        <label className={'news-filters-checkbox'}>
                                            <span>1000 чего-то там</span>
                                            <input type="checkbox"/>
                                            <span className={'news-filters-checkbox_box'}></span>
                                        </label>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </aside>
                    <main className={'b br20 bgc-1 news-main h100'}>

                    </main>
                </div>
            </div>
        </>
    )
}