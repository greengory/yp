import React from 'react';
import ReactStars from 'react-stars';

export const NearbyBusinesses = () => {
    return(
        <div className="businesses-closeby bg-white">
            <div className="container mx-auto py-4 px-8">
                <h3 className="featured_new-businesses text-center py-8">Hot and New Businesses</h3>
                <div className="nearby__business-boxes flex py-8 w-full">
                    <div className="box border-solid bg-black text-black w-1/4 ml-8 mr-8 rounded-sm">
                        <img className="rounded-sm overflow-hidden" src="https://s3-media1.fl.yelpcdn.com/bphoto/5KdmLQivOi4bZj8-COZdKg/l.jpg" />
                        <div className="business__details px-6 py-4">
                            <div>
                                <a className="buziness_name" href="/biz/shushi-factory">Sushi Factory</a>
                                <div className="business_rating flex">
                                    <ReactStars count={5} size={15} color={'#d32323'}  />
                                    <span className="ml-4">5 reviews</span>  
                                </div>
                                <div className="business_categories">
                                    $$  Sushi Bars, Ramen, Poke
                                </div>

                                <div className="business_address">
                                    Hayes Valley
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="box border-solid bg-black text-black w-1/4 mr-8 rounded-sm">
                        <img className="rounded-sm overflow-hidden" src="https://s3-media3.fl.yelpcdn.com/bphoto/z73KNhiYjHW8DrHqILtVZA/l.jpg" />
                        <div className="business__details px-6 py-4">
                            <div>
                                <a className="buziness_name" href="/biz/shushi-factory">Sushi Factory</a>
                                <div className="business_rating flex">
                                    <ReactStars count={5} size={15} color={'#d32323'}  />
                                    <span>5 reviews</span>  
                                </div>
                                <div className="business_categories">
                                    $$  Sushi Bars, Ramen, Poke
                                </div>

                                <div className="business_address">
                                    Hayes Valley
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="box border-solid bg-black text-black w-1/4 mr-8 rounded-sm">
                        <img className="rounded-sm overflow-hidden" src="https://s3-media2.fl.yelpcdn.com/bphoto/Qnu5ywfBniDSdWiOO5SKOg/l.jpg" />
                        <div className="business__details px-6 py-4">
                            <div>
                                <a className="buziness_name" href="/biz/shushi-factory">Sushi Factory</a>
                                <div className="business_rating flex">
                                    <ReactStars count={5} size={15} color={'#d32323'}  />
                                    <span>5 reviews</span>  
                                </div>
                                <div className="business_categories">
                                    $$  Sushi Bars, Ramen, Poke
                                </div>

                                <div className="business_address">
                                    Hayes Valley
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <a className="see_more_businesses text-center underline" href="/business/more_new">See more featured business</a>
            </div>
        </div>
    )
}