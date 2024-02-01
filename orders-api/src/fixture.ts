import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { getDataSourceToken } from '@nestjs/typeorm';
import { DataSource } from 'typeorm';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  await app.init();

  const dataSource = app.get<DataSource>(getDataSourceToken());
  await dataSource.synchronize(true);

  const productRepository = dataSource.getRepository('Product');

  await productRepository.insert([
    {
      id: '1a2b3c4d-5e6f-7g8h-9i0j',
      name: 'product 1',
      description: 'description 1',
      price: 100,
      image_url: 'https://example.com/image1.jpg',
    },
    {
      id: '9i8h7g6f-5e4d-3c2b-1a0j',
      name: 'product 2',
      description: 'description 2',
      price: 200,
      image_url: 'https://example.com/image2.jpg',
    },
    {
      id: '0a1b2c3d-4e5f-6g7h-8i9j',
      name: 'product 3',
      description: 'description 3',
      price: 300,
      image_url: 'https://example.com/image3.jpg',
    },
    {
      id: '1j0i9h8g-7f6e-5d4c-3b2a',
      name: 'product 4',
      description: 'description 4',
      price: 400,
      image_url: 'https://example.com/image4.jpg',
    },
    {
      id: '2a3b4c5d-6e7f-8g9h-0i1j',
      name: 'product 5',
      description: 'description 5',
      price: 500,
      image_url: 'https://example.com/image5.jpg',
    },
    {
      id: '3j2i1h0g-9f8e-7d6c-5b4a',
      name: 'product 6',
      description: 'description 6',
      price: 600,
      image_url: 'https://example.com/image6.jpg',
    },
    {
      id: '4a5b6c7d-8e9f-0g1h-2i3j',
      name: 'product 7',
      description: 'description 7',
      price: 700,
      image_url: 'https://example.com/image7.jpg',
    },
    {
      id: '5j4i3h2g-1f0e-9d8c-7b6a',
      name: 'product 8',
      description: 'description 8',
      price: 800,
      image_url: 'https://example.com/image8.jpg',
    },
    {
      id: '6a7b8c9d-0e1f-2g3h-4i5j',
      name: 'product 9',
      description: 'description 9',
      price: 900,
      image_url: 'https://example.com/image9.jpg',
    },
    {
      id: '7j6i5h4g-3f2e-1d0c-9b8a',
      name: 'product 10',
      description: 'description 10',
      price: 1000,
      image_url: 'https://example.com/image10.jpg',
    },
  ]);
  await app.close();
}
bootstrap();
