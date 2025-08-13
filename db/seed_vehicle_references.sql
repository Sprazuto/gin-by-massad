-- Create default admin user to satisfy foreign key constraints
INSERT INTO
    "user" (
        email,
        password,
        name,
        created_at,
        updated_at
    )
VALUES (
        'admin@eoffice.com',
        '$2a$10$08dUcQTIha4fXdY0CmBWLOp9qpzdR3Jls.bfcdC4k5HECusz4ZcmG',
        'Admin',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    );

-- Wheel Configurations
INSERT INTO
    vehicle_wheels (
        count,
        description,
        created_at,
        updated_at
    )
VALUES (
        2,
        'Roda dua (motor)',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        3,
        'Becak motor/truk kecil',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        4,
        'Standar mobil penumpang',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        6,
        'Truk medium',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        8,
        'Truk besar',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        10,
        'Truk trailer',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    );

-- Popular Vehicle Models in Indonesia
INSERT INTO
    vehicle_model (
        name,
        description,
        created_at,
        updated_at
    )
VALUES (
        'Sedan',
        'Mobil penumpang standar',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'SUV',
        'Kendaraan utilitas sport',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Truck',
        'Pengangkut barang komersial',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Hatchback',
        'Mobil kecil dengan pintu belakang',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Coupe',
        'Mobil sport',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Convertible',
        'Mobil atap terbuka',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Minivan',
        'Kendaraan keluarga',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'MPV',
        'Multi Purpose Vehicle - kendaraan serbaguna',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'LCGC',
        'Low Cost Green Car - mobil murah ramah lingkungan',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Sepeda Motor',
        'Kendaraan roda dua',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Skuter',
        'Motor matik',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Becak Motor',
        'Kendaraan roda tiga',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Truk Pickup',
        'Truk kecil angkutan barang',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Truk Tangki',
        'Kendaraan pengangkut cairan',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Truk Kontainer',
        'Kendaraan pengangkut kontainer',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Truk Tronton',
        'Truk dengan 3 sumbu roda',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Bus',
        'Kendaraan angkut massal',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Mobil Derek',
        'Kendaraan pengangkut mobil rusak',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Mobil Pemadam',
        'Kendaraan pemadam kebakaran',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Ambulans',
        'Kendaraan medis darurat',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Alat Berat',
        'Excavator, bulldozer, dll',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    );

-- Vehicle Colors
INSERT INTO
    vehicle_color (
        name,
        hex_code,
        created_at,
        updated_at
    )
VALUES (
        'Putih',
        '#FFFFFF',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Hitam',
        '#000000',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Silver',
        '#C0C0C0',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Abu-abu',
        '#808080',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Merah',
        '#FF0000',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Biru',
        '#0000FF',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Hijau',
        '#008000',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Kuning',
        '#FFFF00',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Coklat',
        '#A52A2A',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Orange',
        '#FFA500',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Ungu',
        '#800080',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    );

-- Fuel Types
INSERT INTO
    vehicle_fuel (
        type,
        description,
        created_at,
        updated_at
    )
VALUES (
        'Bensin RON 88',
        'Premium - Bensin dengan oktan 88',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Bensin RON 90',
        'Pertalite - Bensin dengan oktan 90',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Bensin RON 92',
        'Pertamax - Bensin dengan oktan 92',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Bensin RON 98',
        'Pertamax Turbo - Bensin dengan oktan 98',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Solar',
        'Biosolar (B20)',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Listrik',
        'Kendaraan listrik (EV)',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Hybrid',
        'Kendaraan hybrid bensin-listrik',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    );

-- Vehicle Ownership Types
INSERT INTO
    vehicle_owning (
        type,
        description,
        created_at,
        updated_at
    )
VALUES (
        'Pemerintahan',
        'Kendaraan dinas',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Perusahaan',
        'Kendaraan perusahaan',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Pribadi',
        'Kendaraan pribadi',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Sewa',
        'Kendaraan disewa',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Kontrak',
        'Kendaraan kontrak operasi',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Pinjam',
        'Kendaraan pinjaman',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    );

-- Vehicle Brands (popular in Indonesia)
INSERT INTO
    vehicle_brand (
        name,
        country,
        founded,
        created_at,
        updated_at
    )
VALUES (
        'Toyota',
        'Jepang',
        1937,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Honda',
        'Jepang',
        1948,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Suzuki',
        'Jepang',
        1909,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Mitsubishi',
        'Jepang',
        1917,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Daihatsu',
        'Jepang',
        1907,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Hyundai',
        'Korea Selatan',
        1967,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'KIA',
        'Korea Selatan',
        1944,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Wuling',
        'China',
        2007,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'DFSK',
        'China',
        1988,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Lexus',
        'Jepang',
        1989,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Isuzu',
        'Jepang',
        1916,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Ford',
        'Amerika Serikat',
        1903,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Chevrolet',
        'Amerika Serikat',
        1911,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Nissan',
        'Jepang',
        1933,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Yamaha',
        'Jepang',
        1955,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Wahana',
        'Indonesia',
        1971,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Kaisar',
        'Indonesia',
        1993,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Tesla',
        'Amerika Serikat',
        2003,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'BYD',
        'China',
        1995,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Hyundai Ioniq',
        'Korea Selatan',
        2016,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Wuling EV',
        'China',
        2020,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'BMW',
        'Jerman',
        1916,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Mercedes-Benz',
        'Jerman',
        1926,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Audi',
        'Jerman',
        1909,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Volkswagen',
        'Jerman',
        1937,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Volvo',
        'Swedia',
        1927,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Kawasaki',
        'Jepang',
        1896,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'MG',
        'China',
        1924,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Jetour',
        'China',
        2018,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Hino',
        'Jepang',
        1942,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Scania',
        'Swedia',
        1891,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'IVECO',
        'Italia',
        1975,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'MAN',
        'Jerman',
        1758,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Tata',
        'India',
        1945,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'UD Trucks',
        'Jepang',
        1935,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Mercedes-Benz Bus',
        'Jerman',
        1895,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Volvo Bus',
        'Swedia',
        1928,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Higer',
        'China',
        1998,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Zhongtong',
        'China',
        1958,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Caterpillar',
        'Amerika Serikat',
        1925,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Komatsu',
        'Jepang',
        1921,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Hitachi',
        'Jepang',
        1910,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Vespa',
        'Italia',
        1946,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Piaggio',
        'Italia',
        1884,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Sym',
        'Taiwan',
        1954,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Benelli',
        'Italia',
        1911,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Federal',
        'Indonesia',
        1984,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Kanzen',
        'Indonesia',
        1986,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Neta',
        'China',
        2018,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'Lexus EV',
        'Jepang',
        1989,
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    );

-- Utilization Companies
INSERT INTO
    utilization_company (
        name,
        address,
        contact,
        created_at,
        updated_at
    )
VALUES (
        'PT Logistik Indonesia',
        'Jl. Raya Industri No. 123, Jakarta',
        '021-1234567',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'PT Transportasi Makmur',
        'Jl. Perkantoran Mega No. 88, Bandung',
        '022-7654321',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'PT Angkasa Jaya',
        'Jl. Gatot Subroto No. 42, Surabaya',
        '031-11223344',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'PT Mitra Trans',
        'Jl. Sudirman No. 25, Medan',
        '061-55667788',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    ),
    (
        'PT Sejahtera Abadi',
        'Jl. Thamrin No. 77, Semarang',
        '024-99887766',
        EXTRACT(
            EPOCH
            FROM NOW()
        ),
        EXTRACT(
            EPOCH
            FROM NOW()
        )
    );