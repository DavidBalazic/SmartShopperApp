import React, { useState } from 'react';
import { ImageIcon } from 'lucide-react';

const ProductImage = ({ src, alt, className = "" }) => {
    const [imageError, setImageError] = useState(false);
    const [imageLoading, setImageLoading] = useState(true);

    return (
        <div className={`relative bg-gray-100 rounded-lg overflow-hidden ${className}`}>
            {imageLoading && (
                <div className="absolute inset-0 flex items-center justify-center">
                    <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-gray-400"></div>
                </div>
            )}
            {!imageError && src ? (
                <img
                    src={src}
                    alt={alt}
                    className="w-full h-full object-cover"
                    onLoad={() => setImageLoading(false)}
                    onError={() => {
                        setImageError(true);
                        setImageLoading(false);
                    }}
                />
            ) : (
                <div className="w-full h-full flex items-center justify-center text-gray-400">
                    <ImageIcon size={24} />
                </div>
            )}
        </div>
    );
};

export default ProductImage;