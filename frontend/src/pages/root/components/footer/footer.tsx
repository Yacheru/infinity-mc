import './footer.css'
import '../../../pages.css'
import React from 'react'

export default function Footer() {
    return (
        <footer className="footer bgc-1 b br20 w100">
            <div className="footer__main h100">
                <div className="footer__items">
                    <a className="footer__item discord" href="https://discord.gg/infinity-tm-494212272353181726"
                       target='_blank'>
                        <svg width="24px" height="24px" viewBox="0 0 24 24" fill="none"
                             xmlns="http://www.w3.org/2000/svg">
                            <path d="M20.317 4.4921C18.7873 3.80147 17.147 3.29265 15.4319 3.00122C15.4007 2.9956 15.3695 3.00965 15.3534 3.03777C15.1424 3.40697 14.9087 3.88862 14.7451 4.26719C12.9004 3.99545 11.0652 3.99545 9.25832 4.26719C9.09465 3.8802 8.85248 3.40697 8.64057 3.03777C8.62449 3.01059 8.59328 2.99654 8.56205 3.00122C6.84791 3.29172 5.20756 3.80054 3.67693 4.4921C3.66368 4.49772 3.65233 4.5071 3.64479 4.51928C0.533392 9.09311 -0.31895 13.5545 0.0991801 17.9606C0.101072 17.9822 0.11337 18.0028 0.130398 18.0159C2.18321 19.4993 4.17171 20.3998 6.12328 20.9967C6.15451 21.0061 6.18761 20.9949 6.20748 20.9695C6.66913 20.3492 7.08064 19.6952 7.43348 19.0073C7.4543 18.967 7.43442 18.9192 7.39186 18.9033C6.73913 18.6597 6.1176 18.3626 5.51973 18.0253C5.47244 17.9981 5.46865 17.9316 5.51216 17.8997C5.63797 17.8069 5.76382 17.7104 5.88396 17.613C5.90569 17.5952 5.93598 17.5914 5.96153 17.6026C9.88928 19.3672 14.1415 19.3672 18.023 17.6026C18.0485 17.5905 18.0788 17.5942 18.1015 17.612C18.2216 17.7095 18.3475 17.8069 18.4742 17.8997C18.5177 17.9316 18.5149 17.9981 18.4676 18.0253C17.8697 18.3692 17.2482 18.6597 16.5945 18.9024C16.552 18.9183 16.533 18.967 16.5538 19.0073C16.9143 19.6942 17.3258 20.3483 17.7789 20.9686C17.7978 20.9949 17.8319 21.0061 17.8631 20.9967C19.8241 20.3998 21.8126 19.4993 23.8654 18.0159C23.8834 18.0028 23.8948 17.9831 23.8967 17.9616C24.3971 12.8676 23.0585 8.4428 20.3482 4.52021C20.3416 4.5071 20.3303 4.49772 20.317 4.4921ZM8.02002 15.2778C6.8375 15.2778 5.86313 14.2095 5.86313 12.8976C5.86313 11.5857 6.8186 10.5175 8.02002 10.5175C9.23087 10.5175 10.1958 11.5951 10.1769 12.8976C10.1769 14.2095 9.22141 15.2778 8.02002 15.2778ZM15.9947 15.2778C14.8123 15.2778 13.8379 14.2095 13.8379 12.8976C13.8379 11.5857 14.7933 10.5175 15.9947 10.5175C17.2056 10.5175 18.1705 11.5951 18.1516 12.8976C18.1516 14.2095 17.2056 15.2778 15.9947 15.2778Z" fill="#000"/>
                        </svg>
                    </a>
                    <a className="footer__item telegram" href="https://t.me/+tpDkWjCoy2BjYjg6" target='_blank'>
                        <svg width="24px" height="24px" viewBox="0 0 48 48" fill="none"
                             xmlns="http://www.w3.org/2000/svg">
                            <path d="M41.4193 7.30899C41.4193 7.30899 45.3046 5.79399 44.9808 9.47328C44.8729 10.9883 43.9016 16.2908 43.1461 22.0262L40.5559 39.0159C40.5559 39.0159 40.3401 41.5048 38.3974 41.9377C36.4547 42.3705 33.5408 40.4227 33.0011 39.9898C32.5694 39.6652 24.9068 34.7955 22.2086 32.4148C21.4531 31.7655 20.5897 30.4669 22.3165 28.9519L33.6487 18.1305C34.9438 16.8319 36.2389 13.8019 30.8426 17.4812L15.7331 27.7616C15.7331 27.7616 14.0063 28.8437 10.7686 27.8698L3.75342 25.7055C3.75342 25.7055 1.16321 24.0823 5.58815 22.459C16.3807 17.3729 29.6555 12.1786 41.4193 7.30899Z" fill="#000000"/>
                        </svg>
                    </a>
                    <a className="footer__item vk" href="https://vk.com/infinitytm" target='_blank'>
                        <svg width="24px" height="24px" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                            <path d="m23.456 5.784c-.27.849-.634 1.588-1.09 2.259l.019-.03q-.672 1.12-1.605 2.588-.8 1.159-.847 1.2c-.138.173-.234.385-.267.618l-.001.007c.027.212.125.397.268.535l.4.446q3.21 3.299 3.611 4.548c.035.092.055.198.055.309 0 .194-.062.373-.167.52l.002-.003c-.176.181-.422.293-.694.293-.03 0-.061-.001-.09-.004h.004-2.631c-.001 0-.003 0-.005 0-.337 0-.647-.118-.89-.314l.003.002c-.354-.291-.669-.606-.951-.948l-.009-.012q-.691-.781-1.226-1.315-1.782-1.694-2.63-1.694c-.021-.002-.045-.003-.07-.003-.165 0-.319.051-.446.138l.003-.002c-.104.13-.167.298-.167.479 0 .036.002.07.007.105v-.004c-.027.314-.043.679-.043 1.048 0 .119.002.237.005.355v-.017 1.159c.01.047.016.101.016.156 0 .242-.11.458-.282.601l-.001.001c-.387.177-.839.281-1.316.281-.102 0-.202-.005-.301-.014l.013.001c-1.574-.03-3.034-.491-4.275-1.268l.035.02c-1.511-.918-2.763-2.113-3.717-3.525l-.027-.042c-.906-1.202-1.751-2.56-2.471-3.992l-.07-.154c-.421-.802-.857-1.788-1.233-2.802l-.06-.185c-.153-.456-.264-.986-.31-1.535l-.002-.025q0-.758.892-.758h2.63c.024-.002.052-.003.081-.003.248 0 .477.085.658.228l-.002-.002c.2.219.348.488.421.788l.003.012c.484 1.367.997 2.515 1.587 3.615l-.067-.137c.482.97 1.015 1.805 1.623 2.576l-.023-.031q.8.982 1.248.982c.009.001.02.001.032.001.148 0 .277-.08.347-.2l.001-.002c.074-.19.117-.411.117-.641 0-.049-.002-.098-.006-.146v.006-3.879c-.021-.457-.133-.884-.32-1.267l.008.019c-.124-.264-.273-.492-.45-.695l.003.004c-.164-.164-.276-.379-.311-.619l-.001-.006c0-.17.078-.323.2-.423l.001-.001c.121-.111.283-.178.46-.178h.008 4.146c.022-.003.047-.004.073-.004.195 0 .37.088.486.226l.001.001c.103.188.164.413.164.651 0 .038-.002.075-.005.112v-.005 5.173c-.002.024-.003.052-.003.08 0 .184.051.357.139.504l-.002-.004c.073.108.195.178.333.178h.001c.176-.012.336-.07.471-.162l-.003.002c.272-.187.506-.4.709-.641l.004-.005c.607-.686 1.167-1.444 1.655-2.25l.039-.07c.344-.57.716-1.272 1.053-1.993l.062-.147.446-.892c.155-.446.571-.76 1.06-.76.019 0 .038 0 .057.001h-.003 2.631q1.066 0 .8.981z" fill="#000000"/>
                        </svg>
                    </a>
                    <a className="footer__item website" href="https://infinity-tm.ru" target='_blank'>
                        <svg width="24px" height="24px" viewBox="0 0 24 24" role="img"
                             xmlns="http://www.w3.org/2000/svg">
                            <path d="M9.103.435a1.753 1.753 0 0 1 1.652-.362c.218.072.406.203.609.333.16.101.348.145.493.261.072.058.014.16.014.232.189.45.29.942.13 1.42-.13.16-.362.203-.55.276-.03.203.043.39.072.594-.043.029-.072.058-.116.087.276-.015.551-.073.827-.13.101-.102.26-.044.39-.059.015-.203.088-.391.088-.594a.34.34 0 0 0 .116-.029c.014.145 0 .29.043.435.073.058.189.029.276.043 0 .058 0 .116.014.174 1.681-.014 3.377 0 5.058 0v.247h.16V2.42h.188c0 .26-.014.507 0 .768a.11.11 0 0 0 .073.029c0 .029.014.087.014.116.058-.058.13-.102.218-.073.014.044.029.087.043.145-.058.058-.087.13-.058.218.464.014.928 0 1.406 0 .044-.058.116-.087.189-.116a.464.464 0 0 1 .087.058h.68a.874.874 0 0 1 .015.333h-.696a.46.46 0 0 1-.087.073c-.072-.044-.13-.073-.202-.116h-1.261c-.247.072-.508.058-.754.014v.275H16.16a.532.532 0 0 1-.29.13c.044.218-.202.276-.29.436a.45.45 0 0 1-.231.174c-.073.449.087.87.203 1.29-.13.029-.276.072-.406.101-.073.29-.145.594-.203.884a.848.848 0 0 1-.507.58c-.174.203-.406.406-.682.42-.101.03-.174-.043-.246-.101-.362.029-.696-.145-1.015-.29-.347-.16-.666-.363-1-.55.03.202-.072.376-.145.55.16.072.377.13.464.304.058.13.116.276.116.435-.014.522-.072 1.044-.101 1.565.014.377-.174.754-.435 1.03-.174-.015-.319-.088-.478-.16-.058.13-.189.26-.116.42.058.189.058.392.145.566a6.674 6.674 0 0 1 1.217 1.753c.304.624.536 1.276.783 1.928.043 0 .13-.015.174-.015.058.189 0 .406.116.58.101.16.072.348.072.536-.029.435-.058.87-.101 1.305-.03.304-.102.609-.145.913.014.232.116.464.101.696-.014.217-.014.449-.188.608.014.493-.116.986.058 1.464.232.32.493.623.768.899.304.145.667.174.928.435.1.16.043.347.014.521a6.732 6.732 0 0 1-1.87 0c-.246-.058-.478-.159-.724-.188-.334.014-.725.145-1.03-.087-.028-.391.117-.768.189-1.145.029-.13.16-.217.145-.348-.03-.45-.058-.913-.087-1.362-.058-.03-.16-.058-.145-.145 0-.218-.072-.435-.13-.638a10.821 10.821 0 0 1-.16-1.681c-.014-.16.087-.276.203-.377.03-.246.058-.507.073-.754-.044-.13-.145-.232-.203-.348-.261.03-.638.087-.797-.188-.377-.565-.769-1.145-1.145-1.71-.16-.015-.348 0-.493-.102-.16-.174-.261-.405-.363-.623-.043.174-.072.362-.174.507-.087.145-.231.247-.333.377-.101.232-.188.464-.275.696-.102.29-.247.58-.276.899a1.543 1.543 0 0 1-.101.449c-.073.116-.203.16-.319.217-.087.189-.145.377-.275.537-.087.101-.232.13-.29.26-.058.174-.145.334-.218.493-.029.174.087.363.03.536-.102.435-.32.841-.522 1.232-.102.29-.174.595-.32.87-.057.116-.202.145-.318.16-.13.318-.276.623-.362.956a5.437 5.437 0 0 0-.03.971c0 .145.088.261.146.391.057.174.014.348-.015.522-.565.073-1.145.13-1.696-.043-.058-.044-.043-.116-.058-.174-.043-.261-.072-.536.015-.783.188-.681.348-1.362.536-2.043-.072-.073-.174-.13-.174-.247-.014-.188 0-.391.044-.58.087-.319.318-.565.434-.87.044-.13.03-.26.044-.39 0-.305.174-.551.304-.812.13-.218.232-.45.406-.638.116-.101.116-.26.203-.391.087-.16.232-.29.232-.479.029-.231-.058-.463-.03-.695.059-.681.19-1.348.305-2.03-.058-.072-.145-.144-.174-.246.015-.072.03-.13.044-.203l-.13-.217c.057-.087.115-.188.173-.275-.058-.044-.13-.102-.188-.145.072-.218.043-.537.304-.638.03.014.102.029.13.043-.043-.376-.043-.768-.086-1.145a3.534 3.534 0 0 1-.073-1.232c.145-.217.42-.304.667-.318-.319-.073-.638-.102-.942-.203-.015-.261.029-.522.072-.783.13-.507.073-1.029.145-1.55.073-.146.261-.175.42-.146.189.03.377-.029.566-.087 0-.072.014-.16 0-.232-.116-.507-.087-1.029 0-1.521.116-.638.377-1.261.855-1.71.319-.305.783-.45 1.217-.435.145 0 .232.145.348.232.058-.058.116-.116.16-.174-.073-.247-.189-.508-.174-.769.029-.58.217-1.174.652-1.565m4.522 4.102c.029.029.029.029 0 0m.203.029c.014.101.043.203-.015.29-.072.029-.16.029-.232.072.203 0 .406.015.61 0 .159-.043.1-.246.1-.362-.1-.116-.318-.044-.463 0m-.507.609c.145.159.217.405.347.565.189-.247.334-.508.551-.725-.26-.015-.522.015-.782-.015-.044.058-.073.116-.116.174z"/>
                        </svg>
                    </a>
                </div>
                <a className='delevoper' href='https://www.donationalerts.com/r/yacheru' target='_blank'>developed by <span className={'important'}>yacheru</span></a>
            </div>
        </footer>
    )
}